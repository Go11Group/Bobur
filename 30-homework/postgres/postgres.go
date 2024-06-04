package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "worker"
	password = "0509"
)

func ConnDB() (*sql.DB, error){
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, user, dbname, password)

	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return	db, nil 
}