package main

import (
	"project/api"
	"project/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	server := api.Routes(db)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
