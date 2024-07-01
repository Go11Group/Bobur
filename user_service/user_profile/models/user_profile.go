package models

type UserRole string


type UserProfile struct {
	Id            string
	UserId        string
	Fullname      string
	Bio           string
	Role          UserRole
	Location      string
	AvatarUrl     string
	Website       string
	Limit, Offset int
}
