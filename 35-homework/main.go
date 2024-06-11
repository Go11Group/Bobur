package main

import (
	"model/handler"
	"model/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	course := postgres.NewCourseRepo(db)
	user := postgres.NewUserRepo(db)

	
	h := handler.NewHandler(db, course, user)
	h.NewServer()
}