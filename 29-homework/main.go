package main

import (
	"model/model"
	"model/strogen/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	// user := mod
	db.AutoMigrate(&model.User11{})

}
