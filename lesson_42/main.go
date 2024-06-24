package main

func main() {

	// Get by id

	// rep, err := http.Get("http://localhost:8080/user/c1b33966-2931-4268-b92a-bd0d55d76c15")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rep.Body.Close()
	// body, _ := io.ReadAll(rep.Body)
	// fmt.Println(string(body))

	// Get all

	// rep, err = http.Get("http://localhost:8080/users")
	// if err != nil {
	// 	panic(err)
	// }

	// body, _ = io.ReadAll(rep.Body)
	// fmt.Println(string(body))

	// CREATE

	// _, err = http.Post("http://localhost:8080/user", "json",
	// 	bytes.NewBuffer([]byte(`{"name":"boburoo1", "email":"aafasf", "birthday":"2005-01-10"}`)))
	// if err != nil {
	// 	panic(err)
	// }

	// Update by id

	// person := http.Client{}

	// req, err := http.NewRequest("PUT", "http://localhost:8080/user/c1b33966-2931-4268-b92a-bd0d55d76c15",
	// 	bytes.NewBuffer([]byte(`{"name":"bobur1", "email":"rrr", "birthday":"2005-10-01"}`)))

	// if err != nil {
	// 	panic(err)
	// }

	// req.Header.Set("Content-Type", "application/json")

	// resp, err := person.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// Delete by id

	// person := http.Client{}
	// req, err := http.NewRequest("DELETE", "http://localhost:8080/user/c1b33966-2931-4268-b92a-bd0d55d76c15", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := person.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

}
