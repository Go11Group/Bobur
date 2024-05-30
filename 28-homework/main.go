package main

import (
	"conn/asd"
	"conn/storge/posgres"
	"fmt"
)

func main() {
	db, err := posgres.ConnectDB()
	if err != nil {
		fmt.Println(333)
		panic(err)
	}
	defer db.Close()

	course := posgres.NewCourseRepo(db)
	student := posgres.NewStudentRepo(db)
	// c, err := course.GetAllCourses()
	// if err != nil {
	// 	panic(err)
	// }

	for {
		fmt.Println("0.Exit\t1.Course\t2.Student")
		var a int
		fmt.Scan(&a)
		switch a {
		case 0:
			fmt.Println("Successfully connected!")
			return
		case 1:
			asd.CourseCRDU(course)
		case 2:
			asd.StudentCRDU(student)
		}
	}

}
