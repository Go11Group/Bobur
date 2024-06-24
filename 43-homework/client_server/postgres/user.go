package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"model/model"

	"github.com/google/uuid"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// Creat user
func (u *UserRepo) CreateUser(user *model.User) error {
	user.Id = uuid.NewString()
	_, err := u.Db.Exec("insert into user2(id, name, phone, age) values($1, $2, $3, $4)",
		user.Id, user.Name, user.Phone, user.Age)
	if err != nil {
		return err
	}
	return nil
}

// Update user by id and body
func (u *UserRepo) UpdateUser(id string, user *model.User) (*model.User, error) {
	_, err := u.Db.Exec("update user2 set name=$1, phone=$2, age=$3 where id=$4",
		user.Name, user.Phone, user.Age, id)

	if err != nil {
		return nil, err
	}
	user1, err := u.GetById(id)
	if err != nil {
		return nil, err
	}

	return user1, nil
}

// Delete user by id
func (u *UserRepo) DeleteUser(id string) error {
	_, err := u.Db.Exec("delete from user2 where id=$1", id)
	if err != nil {
		return err
	}

	return err
}

// Get by id
func (u *UserRepo) GetById(id string) (*model.User, error) {
	user := model.User{}
	row := u.Db.QueryRow("select id, name, phone, age from user2 where id=$1", id)
	err := row.Scan(&user.Id, &user.Name, &user.Phone, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %s not found", id) // maxsus xatolik
		}
		return nil, err // boshqa xatoliklar
	}
	return &user, nil
}

// Get all
func (u *UserRepo) GetByAll() ([]model.User, error) {
	rows, err := u.Db.Query("SELECT id, name, age, phone FROM user2")
	if err != nil {
		log.Println("Query failed:", err)
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}

	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Phone)
		if err != nil {
			log.Println("Failed to scan row:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Println("Rows iteration failed:", err)
		return nil, err
	}

	return users, nil
}
