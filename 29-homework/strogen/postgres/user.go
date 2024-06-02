package postgres

import (
	"model/model"

	"gorm.io/gorm"
)

type User11Repo struct {
	DB *gorm.DB
}

func NewUser11Repo(DB *gorm.DB) *User11Repo {
	return &User11Repo{DB}
}

func (u *User11Repo) CreateUser11(user *model.User11) error {
	result := u.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *User11Repo) UpdateUser11(user model.User11) error {
	result := u.DB.Model(&model.User11{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"Name": user.LastName,
		"Age":  user.Age,
		"Email": user.Email,
		"Password": user.Password,
	})
	return result.Error
}

