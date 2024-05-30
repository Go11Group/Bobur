package posgres

import (
	"conn/model"
	"database/sql"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) CreateCourse(name string) error {
	_, err := c.DB.Exec("insert into course(course_name) values($1)", &name)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *CourseRepo) DeleteCourse(id string) error {
	_, err := c.DB.Exec("DELETE FROM course WHERE course_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo) UpdateCourse(name, id string) error {
    _, err := c.DB.Exec("UPDATE course SET course_name = $1 WHERE course_id = $2", name, id)
    if err != nil {
        return err
    }
    return nil
}


func (c *CourseRepo) GetAllCourses() ([]model.Course, error) {
	rows, err := c.DB.Query("select course_name, course_id from course")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var courses []model.Course

	for rows.Next() {
		var course model.Course
		err = rows.Scan(&course.Name, &course.Id)
		if err != nil {
			panic(err)
		}
		courses = append(courses, course)

	}
	return courses, nil

}

func (c *CourseRepo) GetCourseId(id string) (model.Course, error) {
	var course model.Course
	c.DB.QueryRow("select course_id, course_name from course where course_id = $1", id).Scan(&course.Id, &course.Name)

	return course, nil
}
