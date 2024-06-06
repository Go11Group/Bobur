package main

import (
	"fmt"
	"net/http"
	"strings"
)

type person struct {
	Name     string
	LasrName string
	Age      string
}

var users = map[string]string{"1": "bobur", "2": "jamshid", "3": "husan"}

func main() {
	http.HandleFunc("GET /user/", user)
	http.HandleFunc("/person", person)

	err := http.ListenAndServe(":7075", nil)
	if err != nil {
		panic(err)
	}

}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Write([]byte(users[strings.TrimPrefix(r.URL.Path, "/user/")]))
}


func person(w http.ResponseWriter, r *http.Request)
