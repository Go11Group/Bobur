package asd

import (
	"conn/model"
	"conn/storge/posgres"
	"fmt"
)

func CourseCRDU(course *posgres.CourseRepo) {
	fmt.Println("0.Exit\t1.Create Course\t2.Delete Course\t3.Update Course\t4.Read Course\t5.Issue by id")
	var b int
	fmt.Scan(&b)
	switch b {
	case 0:
		break
	case 1:
		var name string
		fmt.Printf("Enter the name of the course: ")
		fmt.Scan(&name)
		err := course.CreateCourse(name)
		if err != nil {
			panic(err)
		}
	case 2:
		var id string
		fmt.Printf("Enter the id of the course: ")
		fmt.Scan(&id)
		err := course.DeleteCourse(id)
		if err != nil {
			panic(err)
		}
	case 3:
		var id, name string
		fmt.Printf("Enter the id of the course: ")
		fmt.Scan(&id)
		fmt.Printf("Enter a new name for the course: ")
		fmt.Scan(&name)
		err := course.UpdateCourse(name, id)
		if err != nil {
			panic(err)
		}
	case 4:
		var courses []model.Course
		courses, err := course.GetAllCourses()
		if err != nil {
			panic(err)
		}
		for _, k := range courses {
			fmt.Println(k.Id, k.Name)
		}
	case 5:
		var id string
		var course1 model.Course
		fmt.Printf("Enter the id of course: ")
		fmt.Scan(&id)
		course1, err := course.GetCourseId(id)
		if err != nil {
			panic(err)
		}
		fmt.Println(course1)
	}
}
