package postgres

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	pb "portfolio_service/genproto"

	"github.com/google/uuid"
)

type EducationStorage struct {
	db *sql.DB
}

func NewEducationStorage(db *sql.DB) *EducationStorage {
	return &EducationStorage{
		db: db,
	}
}

func (s *EducationStorage) CreateEducation(education *pb.CreateEducationRequest) (*pb.Education, error) {
	education.Id = uuid.New().String()
	query :=
		`
	INSERT INTO education(
		id,
		user_id,
		institution,
		degree,
		field_of_study,
		start_date,
		end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := s.db.Exec(query,
		education.Id,
		education.UserId,
		education.Institution,
		education.Degree,
		education.FieldOfStudy,
		education.StartDate,
		education.EndDate,
	)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *EducationStorage) UpdateEducation(education *pb.UpdateEducationRequest) (*pb.Void, error) {
	query :=
		`
	UPDATE education
	SET
		user_id = $2,
		institution = $3,
		degree = $4,
		field_of_study = $5,
		start_date = $6,
		end_date = $7
	WHERE 
		id = $1 and deleted_at = 0
	`
	_, err := s.db.Exec(query, education.Id,education.UserId, education.Institution, education.Degree, education.FieldOfStudy, education.StartDate, education.GetEndDate())

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return nil, nil
}

func (s *EducationStorage) DeleteEducation(education *pb.ById) (*pb.Void, error) {
	query :=
		`
		update education set deleted_at = EXTRACT(EPOCH FROM NOW()) where id = $1
	`
	_, err := s.db.Exec(query, education.GetId())
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return nil, nil
}

func (s *EducationStorage) GetEducation(id *pb.ById) (*pb.Education, error) {
	query :=
		`
	SELECT
		id,
		user_id,
		institution,
		degree,
		field_of_study,
		start_date,
		end_date
	FROM
		education
	WHERE
		id = $1 and deleted_at = 0
	`
	education := &pb.Education{}

	err := s.db.QueryRow(query, id.Id).Scan(
		&education.Id,
		&education.UserId,
		&education.Institution,
		&education.Degree,
		&education.FieldOfStudy,
		&education.StartDate,
		&education.EndDate,
	)
	if err != nil {
		log.Println("Error querying education record:", err)
		return nil, err
	}
	return education, nil
}

func (s *EducationStorage) GetEducations(flt *pb.EducationFilter) (*pb.Educations, error) {
	mainQuery := `
	SELECT
		id,
		user_id,
		institution,
		degree,
		field_of_study,
		start_date,
		end_date
	FROM
		education
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
	if flt.Institution != "" {
		query = append(query, "institution = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.Institution)
		paramIndex++
	}
	if flt.Degree != "" {
		query = append(query, "degree = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.Degree)
		paramIndex++
	}
	if flt.FieldOfStudy != "" {
		query = append(query, "field_of_study = $"+strconv.Itoa(paramIndex))
		args = append(args, flt.FieldOfStudy)
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
		return nil, err
	}

	defer rows.Close()
	educations := []*pb.Education{}
	for rows.Next() {
		education := &pb.Education{}
		err := rows.Scan(
			&education.Id,
			&education.UserId,
			&education.Institution,
			&education.Degree,
			&education.FieldOfStudy,
			&education.StartDate,
			&education.EndDate,
		)
		if err != nil {
			return nil, err
		}
		educations = append(educations, education)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &pb.Educations{Educations: educations}, nil

}
