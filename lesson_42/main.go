package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// CREATE
	// _, err := http.Post("http://localhost:8080/user", "json",
	// 	bytes.NewBuffer([]byte(`{"name":"boburoo1", "email":"aafasf", "birthday":"2005-01-10"}`)))
	// if err != nil {
	// 	panic(err)
	// }

	// Get by id
	rep, err := http.Get("http://localhost:8080/user/c1b33966-2931-4268-b92a-bd0d55d76c15")
	if err != nil {
		panic(err)
	}
	defer rep.Body.Close()
	body, _ := io.ReadAll(rep.Body)
	fmt.Println(string(body))


}
