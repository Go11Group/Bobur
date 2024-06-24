package postgres

import (
	"database/sql"
	"fmt"
	filters "model/filter"
	"model/model"
	"time"
)

type EnrollmentRepo struct {
	Db *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db}
}


// Yangi Enrollment yaratish 
func (e *EnrollmentRepo) CreateEnrollment(enrollment *model.Enrollments) error {

	_, err := e.Db.Exec("insert into enrollments(user_id, course_id, enrollment_date, create_at)"+
		"values($1, $2, $3, $4)", enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate, time.Now())

	if err != nil {
		return err
	}

	return nil
}

// Enrollment id si orqali Update
func (e *EnrollmentRepo) UpdateEnrollment(enrollment *model.Enrollments, id string) error {
	_, err := e.Db.Exec("update enrollments set enrollment_date=$1, update_at=$2 where enrollment_id=$3",
		enrollment.EnrollmentDate, time.Now(), id)
	if err != nil {
		fmt.Printf("Update Query error: %v", err)
		return err
	}

	return nil
}

// Enrollment id si orqli o`chirish 
func (e *EnrollmentRepo) DeleteEnrollment(id string) error {
	_, err := e.Db.Exec("delete from enrollments where enrollment_id=$1", id)
	if err != nil {
		fmt.Printf("Create Query error: %v", err)
		return err
	}
	return nil
}

// Enrollement tableni columin lar orqali filterlash
func (e *EnrollmentRepo) GetAllFilterEnrollments(f *model.FilterEnrollments) ([]model.Enrollments, error) {
	filter := ""
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)

	query := `select enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at
	 	from enrollments where true `

	if len(f.Limit) > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	if len(f.UserId) > 0 {
		params["user_id"] = f.UserId
		filter += " and user_id = :user_id "
	}

	if len(f.CourseId) > 0 {
		params["course_id"] = f.CourseId
		filter += " and course_id = :course_id "
	}

	if len(f.CreatedAt) > 0 {
		params["created_at"] = f.CreatedAt
		filter += " and created_at = :created_at "
	}

	if len(f.UpdatedAd) > 0 {
		params["updated_at"] = f.UpdatedAd
		filter += " and updated_at = :updated_at "
	}

	if len(f.EnrollmentDate) > 0 {
		params["enrollment_data"] = f.EnrollmentDate
		filter += " and enrollment_data = :enrollment_data "
	}

	query = query + filter + limit

	query, arr = filters.ReplaceQueryParams(query, params)
	rows, err := e.Db.Query(query, arr...)
	if err != nil {
		fmt.Println("111 filterda xato bor course")
		return []model.Enrollments{}, err
	}

	enrollments := []model.Enrollments{}
	for rows.Next() {
		enrollment := model.Enrollments{}
		err = rows.Scan(&enrollment.Id, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
		if err != nil {
			fmt.Println("2222 filtersda xato bor course")
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}

func (e *EnrollmentRepo) GetIdEnrollment(id string) (model.Enrollments, error) {
	enrollment := model.Enrollments{}
	row := e.Db.QueryRow("select enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at from enrollments where enrollment_id=$1", id)

	err := row.Scan(&enrollment.Id, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
	if err != nil {
		return model.Enrollments{}, err
	}
	return enrollment, nil

}
