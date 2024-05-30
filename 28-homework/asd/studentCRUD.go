package asd

import (
	"conn/model"
	"conn/storge/posgres"
	"fmt"
)

func StudentCRDU(student *posgres.StudentRepo) {
	fmt.Println("0.Exit\t1.Create Student\t2.Delete Student\t3.Update Student\t4.Read Student\t5.Issue by id")
	var b int
	fmt.Scan(&b)
	switch b {
	case 0:
		break
	case 1:
		var student1 model.Student
		fmt.Printf("Isim kiriting: ")
		fmt.Scan(&student1.Name)
		fmt.Printf("Familya kiriting: ")
		fmt.Scan(&student1.Lastname)
		fmt.Printf("Telefon raqam kiriting: ")
		fmt.Scan(&student1.Phone)
		fmt.Printf("Yoshingizni kiriting: ")
		fmt.Scan(&student1.Age)
		err := student.CreateStudent(student1)
		if err != nil {
			panic(err)
		}
	case 2:
		var id string
		fmt.Printf("Enter the id of the student: ")
		fmt.Scan(&id)
		err := student.DeleteStudent(id)
		if err != nil {
			panic(err)
		}
	case 3:
		var id, name string
		fmt.Printf("Enter the id of the course: ")
		fmt.Scan(&id)
		fmt.Printf("Enter a new name for the course: ")
		fmt.Scan(&name)
		err := student.UpdateStudent(name, id)
		if err != nil {
			panic(err)
		}
	case 4:
		var students []model.Student
		students, err := student.GetAllStudent()
		if err != nil {
			panic(err)
		}
		for _, k := range students {
			fmt.Println(k.Id, k.Name)
		}
	case 5:
		var id string
		var student1 model.Student
		fmt.Printf("Enter the id of course: ")
		fmt.Scan(&id)
		student1, err := student.GetStudentId(id)
		if err != nil {
			panic(err)
		}
		fmt.Println(student1)
	}
}
