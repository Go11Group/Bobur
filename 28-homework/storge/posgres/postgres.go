package posgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	dbname   = "reja"
	password = "0509"
	port     = 5432
)

func ConnectDB() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println(111)
		panic(err)	
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(222)
		panic(err)
	}

	return db, nil
}
