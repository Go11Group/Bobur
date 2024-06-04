package main

import (
	"fmt"
	"math/rand"
	"model/postgres"

	"github.com/go-faker/faker/v3"
)


func main() {
	db, err := postgres.ConnDB()

	if err != nil {
		panic(err)	
	}
	defer  db.Close()

	for i:=0; i< 1_000_000; i++ {
		_, err = db.Exec("insert into user1(name, fname, phone, salary) values($1, $2, $3, $4)", 
		faker.FirstName(), faker.LastName(), rand.Intn(10000)+20, rand.Intn(20)+100)

		if err != nil {
			panic(err)
		}

		if i % 1000 == 0{
			fmt.Println(i)
		}
	}
}