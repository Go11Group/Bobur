package models

type User struct {
	Id       int    `json:id`
	Name     string `json:name`
	Lastname string `json:lastname`
	Phone    string `json:phone`
}
