package postgres

import (
	"database/sql"
	"strconv"
	"strings"

	pb "portfolio_service/genproto"

	"github.com/google/uuid"
)

type ProjectRepo struct {
	db *sql.DB
}

func NewProjectStorage(db *sql.DB) *ProjectRepo {
	return &ProjectRepo{
		db: db,
	}
}

func (s *ProjectRepo) CreateProject(project *pb.CreateProjectRequest) (*pb.Project, error) {
	project.Id = uuid.New().String()
	query :=
		`
	INSERT INTO projects(
		id,
		user_id,
		title,
		description,
		url
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := s.db.Exec(query,

		project.UserId,
		project.Title,
		project.Description,
		project.Url,
		project.Id,
	)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *ProjectRepo) UpdateProject(project *pb.UpdateProjectRequest) (*pb.Void, error) {
	baseQuery := ` UPDATE projects SET `

	if project.UserId != "" {
		baseQuery += ` user_id = $2, `
	}
	if project.Title != "" {
		baseQuery += ` title = $3, `
	}
	if project.Description != "" {
		baseQuery += ` description = $4, `
	}
	if project.Url != "" {
		baseQuery += ` url = $5, `
	}
	baseQuery = strings.TrimSuffix(baseQuery, ", ")
	baseQuery += ` WHERE id = $1 AND deleted_at = 0 `
	_, err := s.db.Exec(baseQuery,
		project.Id,
		project.UserId,
		project.Title,
		project.Description,
		project.Url,
	)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *ProjectRepo) DeleteProject(project *pb.ById) (*pb.Void, error) {
	query := `UPDATE projects SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`

	_, err := s.db.Exec(query, project.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *ProjectRepo) GetProject(flt *pb.ById) (*pb.Project, error) {
	query := `SELECT id, user_id, title, description, url FROM projects WHERE id = $1 AND deleted_at = 0`

	row := s.db.QueryRow(query, flt.Id)
	res := &pb.Project{}
	err := row.Scan(&res.Id, &res.UserId, &res.Title, &res.Description, &res.Url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProjectRepo) GetProjects(flt *pb.ProjectFilter) (*pb.Projects, error) {

	baseQuery := `SELECT id, user_id, title, description, url FROM projects WHERE deleted_at = 0`
	var cond []string
	var args []interface{}
	paramIndex := 1

	if flt.UserId != "" {
		cond = append(cond, "user_id = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.UserId)
		paramIndex++
	}

	if flt.Url != "" {
		cond = append(cond, "url = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.Url)
		paramIndex++
	}

	if len(cond) > 0 {
		baseQuery += " AND " + strings.Join(cond, " AND ")
	}

	rows, err := s.db.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []*pb.Project
	for rows.Next() {
		res := &pb.Project{}
		err = rows.Scan(&res.Id, &res.UserId, &res.Title, &res.Description, &res.Url)
		if err != nil {
			return nil, err
		}
		projects = append(projects, res)
	}
	return &pb.Projects{Projects: projects}, nil
}
