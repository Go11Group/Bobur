package posgres

import (
	"conn/model"
	"database/sql"
)

type StudentRepo struct {
	DB *sql.DB
}

func NewStudentRepo(DB *sql.DB) *StudentRepo {
	return &StudentRepo{DB}
}

func (s *StudentRepo) CreateStudent(student model.Student) error {
	_, err := s.DB.Exec("insert into student(name, lastname, phone, age) values($1, $2, $3, $4)", &student.Name, &student.Lastname, &student.Phone, &student.Age)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *StudentRepo) DeleteStudent(id string) error {
	_, err := s.DB.Exec("DELETE FROM student WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *StudentRepo) UpdateStudent(name, id string) error {
	_, err := s.DB.Exec("update student set name = $1 where id = $2", name, id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *StudentRepo) GetAllStudent() ([]model.Student, error) {
	rows, err := s.DB.Query("SELECT name, lastname, phone, age FROM student")
	if err != nil {

		panic(err)
	}
	defer rows.Close()

	var students []model.Student

	for rows.Next() {
		var student model.Student
		err = rows.Scan(&student.Name, &student.Lastname, &student.Phone, &student.Age)
		if err != nil {
			panic(err)
		}
		students = append(students, student)

	}
	return students, nil

}

func (s *StudentRepo) GetStudentId(id string) (model.Student, error) {
	var student model.Student
	s.DB.QueryRow("select name, lastname, phone, age from course where id = $1", &id).Scan(&student.Id, &student.Name, &student.Lastname, &student.Phone, &student.Age)

	return student, nil
}
