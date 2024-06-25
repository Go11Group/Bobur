package postgres

import (
	"database/sql"
	"strconv"
	"strings"

	pb "portfolio_service/genproto"

	"github.com/google/uuid"
)

type SkillRepo struct {
	db *sql.DB
}

func NewSkillStorage(db *sql.DB) *SkillRepo {
	return &SkillRepo{
		db: db,
	}
}

func (s *SkillRepo) CreateSkill(skill *pb.CreateSkillRequest) (*pb.Skill, error) {
	skill.Id = uuid.New().String()
	query :=
		`
	INSERT INTO skills(
		id,
		user_id,
		name,
		level)
		VALUES ($1, $2, $3, $4)
	`

	_, err := s.db.Exec(query,
		skill.Id,
		skill.UserId,
		skill.Name,
		skill.Level,
	)
	if err != nil {
		return nil, err
	}

	return nil, nil // Return the created skill
}

func (s *SkillRepo) UpdateSkill(skill *pb.UpdateSkillRequest) (*pb.Void, error) {
	baseQuery := ` UPDATE skills SET `

	if skill.UserId != "" {
		baseQuery += ` user_id = $2, `
	}
	if skill.Name != "" {
		baseQuery += ` name = $3, `
	}
	if skill.Level != "" {
		baseQuery += ` level = $4, `
	}
	baseQuery = strings.TrimSuffix(baseQuery, ", ")
	baseQuery += ` WHERE id = $1 AND deleted_at = 0 `
	_, err := s.db.Exec(baseQuery,
		skill.Id,
		skill.UserId,
		skill.Name,
		skill.Level,
	)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *SkillRepo) DeleteSkill(skill *pb.ById) (*pb.Void, error) {
	query := `UPDATE skills SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err := s.db.Exec(query, skill.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}


func (s *SkillRepo) GetSkill(flt *pb.ById) (*pb.Skill, error) { // Adjust return type to *pb.Skill
	query := `SELECT id, user_id, name, level FROM skills WHERE id = $1 AND deleted_at = 0`
	row := s.db.QueryRow(query, flt.Id)
	res := &pb.Skill{}
	err := row.Scan(&res.Id, &res.UserId, &res.Name, &res.Level)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SkillRepo) GetSkills(flt *pb.SkillFilter) (*pb.Skills, error) {
	baseQuery := ` SELECT id, user_id, name, level FROM skills WHERE deleted_at = 0 `

	var cond []string
	var args []interface{}
	paramIndex := 1
	if flt.UserId != "" {
		cond = append(cond, "user_id = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.UserId)
		paramIndex++
	}

	if flt.Name != "" {
		cond = append(cond, "name = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.Name)
		paramIndex++
	}

	if len(cond) > 0 {
		baseQuery += " AND " + strings.Join(cond, " AND ")
	}
	row, err := s.db.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	res := &pb.Skills{}
	for row.Next() {
		skill := &pb.Skill{}
		err := row.Scan(&skill.Id, &skill.UserId, &skill.Name, &skill.Level)
		if err != nil {
			return nil, err
		}
		res.Skills = append(res.Skills, skill)
	}
	return res, err
}

