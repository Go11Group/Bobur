package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)



const (
	host = "localhost"
	port = 5432
	dbname = "reja"
	password = "0509"
	user = "postgres"
)


func ConnDB() (*sql.DB, error){
	conn := fmt.Sprintf("host=%s port=%d dbname=%s password=%s user=%s sslmode=disable", host, port, dbname, password, user)

	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		 panic(err)
	}

	return db, nil
}