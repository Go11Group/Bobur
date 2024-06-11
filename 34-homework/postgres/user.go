package postgres

import (
	"database/sql"
	"model/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) CreateUser(user models.User) error {
	_, err := u.DB.Exec("insert into user1(name, fname, phone) values($1, $2, $3)", user.Name, user.Lastname, user.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) DeleteUser(id int) error {
	_, err := u.DB.Exec("delete from user1 where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) UpdateUser(user models.User) error {
	_, err := u.DB.Exec("update user1 set name=$1, fname=$2, phone=$3 where id=$4", user.Name, user.Lastname, user.Phone, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetUserId(id int) (models.User, error) {
	row := u.DB.QueryRow("select id, name, fname, phone from user1 where id=$1", id)
	user := models.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Lastname, &user.Phone)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepo) GetAllUser() ([]models.User, error) {
	users := []models.User{}
	rows, err := u.DB.Query("select id, name, fname, phone from user1")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Lastname, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
