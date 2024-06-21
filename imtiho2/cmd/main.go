package main

import (
	"fmt"
	"model/handler"
	"model/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	user := postgres.NewUserRepo(db)
	course := postgres.NewCourseRwpo(db)
	lesson := postgres.NewLessonRwpo(db)
	enrollment := postgres.NewEnrollmentRepo(db)

	handler := handler.NewHandler(db, user, course, lesson, enrollment)
	r := handler.NewServer()
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
