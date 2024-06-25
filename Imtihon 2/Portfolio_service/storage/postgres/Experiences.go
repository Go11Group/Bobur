package postgres

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	pb "portfolio_service/genproto"

	"github.com/google/uuid"
)

type ExperienceRepo struct {
	db *sql.DB
}

func NewExperienceStorage(db *sql.DB) *ExperienceRepo {
	return &ExperienceRepo{
		db: db,
	}
}

func (s *ExperienceRepo) CreateExperience(experience *pb.CreateExperienceRequest) (*pb.Experience, error) {
	experience.Id = uuid.New().String()
	query :=
		`
	INSERT INTO experiences(
		id,
		user_id,
		title,
		company,
		description,
		start_date,
		end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := s.db.Exec(query,
		experience.Id,
		experience.UserId,
		experience.Title,
		experience.Company,
		experience.Description,
		experience.StartDate,
		experience.EndDate,
	)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *ExperienceRepo) UpdateExperience(experience *pb.UpdateExperienceRequest) (*pb.Void, error) {
	query :=

		`
	UPDATE experiences
	SET
		user_id = $2,
		title = $3,
		company = $4,
		description = $5
		
	WHERE 
		id = $1 and deleted_at = 0
	`
	_, err := s.db.Exec(query, experience.Id, experience.GetUserId(), experience.GetTitle(), experience.GetCompany(), experience.GetDescription())

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return nil, nil
}

func (s *ExperienceRepo) DeleteExperience(experience *pb.ById) (*pb.Void, error) {
	query :=
		`
		update experiences set deleted_at = EXTRACT(EPOCH FROM NOW()) where id = $1
	`
	_, err := s.db.Exec(query, experience.GetId())
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return nil, nil
}

func (s *ExperienceRepo) GetExperience(flt *pb.ById) (*pb.Experience, error) {

	mainQuery := `
	SELECT
		id,
		user_id,
		title,
		company,
		description,
		start_date,
		end_date
	FROM
		experiences
	WHERE
		id = $1
		AND deleted_at = 0
	`
	row := s.db.QueryRow(mainQuery, flt.Id)
	res := &pb.Experience{}
	err := row.Scan(&res.Id, &res.UserId, &res.Title, &res.Company, &res.Description, &res.StartDate, &res.EndDate)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ExperienceRepo) GetExperiences(flt *pb.ExperienceFilter) (*pb.Experiences, error) {
	mainQuery := `
	SELECT
		id,
		user_id,
		title,
		company,
		description,
		start_date,
		end_date
	FROM
		experiences
	WHERE
		deleted_at = 0
	`

	var query []string
	var args []interface{}
	paramIndex := 1

	if flt.UserId != "" {
		query = append(query, "user_id = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.UserId)
		paramIndex++
	}

	if flt.Company != "" {
		query = append(query, "company = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.Company)
		paramIndex++
	}

	if flt.StartDate != "" {
		query = append(query, "start_date = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.StartDate)
		paramIndex++
	}

	if flt.EndDate != "" {
		query = append(query, "end_date = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.EndDate)
		paramIndex++
	}

	if len(query) > 0 {
		mainQuery += " AND " + strings.Join(query, " AND ")
	}

	rows, err := s.db.Query(mainQuery, args...)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer rows.Close()

	experiences := &pb.Experiences{}
	for rows.Next() {
		experience := &pb.Experience{}
		err := rows.Scan(
			&experience.Id,
			&experience.UserId,
			&experience.Title,
			&experience.Company,
			&experience.Description,
			&experience.StartDate,
			&experience.EndDate,
		)
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		experiences.Experiences = append(experiences.Experiences, experience)
	}
	return experiences, nil
}
