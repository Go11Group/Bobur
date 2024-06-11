package postgres

import (
	"database/sql"
	"model/models"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db}
}

func (c *CourseRepo) CreateCourse(course models.Course) error {
	_, err := c.DB.Exec("insert into course1(name) values($1)", course.Name)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo) DeleteCourse(id int) error {
	_, err := c.DB.Exec("delete from course1 where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo) UpdateCourse(course models.Course) error {
	_, err := c.DB.Exec("update course1 set name=$1 where id=$2", course.Name, course.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo) GetCourseId(id int) (models.Course, error) {
	row := c.DB.QueryRow("select id, name from course1 where id=$1", id)
	course := models.Course{}
	err := row.Scan(&course.Id, &course.Name)
	if err != nil {
		return course, err
	}
	return course, nil
}

func (c *CourseRepo) GetAllCourse() ([]models.Course, error) {
	courses := []models.Course{}
	rows, err := c.DB.Query("select id, name from course1")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		course := models.Course{}
		err = rows.Scan(&course.Id, &course.Name)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
