package models

type Users struct {
	Id string
	UserName string
	Email string
	Password_hash string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Limit,Offset int 
}