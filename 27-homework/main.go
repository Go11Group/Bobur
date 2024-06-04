package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	Id    string
	Name  string
	Fname string
	Phone string
}

type Course struct {
	Id   int
	Name string
}

func main() {
	user := User{}
	course := Course{}
	connStr := "user=postgres password=0509 dbname=nt sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
	}
	defer db.Close()

	err = db.QueryRow("SELECT id, name, fname, phone FROM user1").Scan(&user.Id, &user.Name, &user.Fname, &user.Phone)
	if err != nil {
		fmt.Println(1)
		panic(err)
	}
	fmt.Println(user)

	err = db.QueryRow("select u.name, c.name from user1 as u "+
		"left join course1 as c on c.id = u.id where length(u.name) > 5").Scan(&user.Name, &course.Name)

	if err != nil {
		fmt.Println(1)
		panic(err)
	}
	fmt.Println("----------------")
	fmt.Println(user.Name, course.Name)
	fmt.Println("----------------")

	// -------------------- Query() ------------------------------------------------------
	rows, err := db.Query("SELECT id, name, fname, phone FROM user1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Fname, &user.Phone)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID NAME  FNAME  PHONE")
	fmt.Println("------------------------------")
	for _, k := range users {
		fmt.Println(k.Id, "|", k.Name, "|", k.Fname, "|", k.Phone)
		fmt.Println("------------------------------")
	}
	// ------------------------------------------------------------------------------------------
	
	

	rows, err = db.Query("SELECT u.name, c.name FROM user1 as u left join course1 as c on c.id = u.id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var uname, cname string
	for rows.Next() {
		err = rows.Scan(&uname, &cname)
		if err != nil {
			panic(err)
		}
		fmt.Println(uname, cname)
	}

	fmt.Println("Muvafaqqiyatli amalga oshirildi")
}
