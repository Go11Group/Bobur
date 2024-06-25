package postgres

import (
	"database/sql"
	"fmt"
	"strconv"

	"portfolio_service/storage"

	"portfolio_service/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	db         *sql.DB
	education  storage.EducationI
	experience storage.ExperienceI
	project    storage.ProjectsI
	skills     storage.SkillsI
}

func ConnectDB() (storage.StorageI, error) {
	cfg := config.Load()
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		strconv.Itoa(cfg.PostgresPort), // Convert PostgresPort from int to string
		cfg.PostgresDatabase,
	)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return &Storage{db: db}, nil
}

func (s *Storage) Education() storage.EducationI {
	if s.education == nil {
		s.education = &EducationStorage{
			db: s.db,}
	}
	return s.education
}

func (s *Storage) Experience() storage.ExperienceI {
	if s.experience == nil {
		s.experience = &ExperienceRepo{
			db: s.db,}
	}
	return s.experience
}

func (s *Storage) Projects() storage.ProjectsI {
	if s.project == nil {
		s.project = &ProjectRepo{
			db: s.db,}
	}
	return s.project
}

func (s *Storage) Skills() storage.SkillsI {
	if s.skills == nil {
		s.skills = &SkillRepo{
			db: s.db,}
	}
	return s.skills
}
