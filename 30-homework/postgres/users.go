package postgres

import (
	"database/sql"
	"model/strac"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) CreateUser(user strac.User) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tr.Exec("insert into users(username, email, password) values($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		tr.Rollback()
		return err
	}

	return tr.Commit()
}

func (u *UserRepo) UpdateUser(user strac.User) error {
	tr, err := u.DB.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tr.Exec("update users set username=$1, email=$2, password=$3 where id = $4", user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		tr.Rollback()
		return err
	}
	return tr.Commit()
}

func (u *UserRepo) DeleteUser(id int) error {
	tr, err := u.DB.Begin()
	if err != nil {
		panic(err)
	}

	_, err = tr.Exec("delete from users where id=$1", id)
	if err != nil {
		tr.Rollback()
		return err
	}
	return tr.Commit()
}

func (u *UserRepo) GetAllUser() ([]strac.User, error) {
	tr, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tr.Query("select username, email, password from users")
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	defer rows.Close()

	var users []strac.User
	for rows.Next() {
		var user strac.User
		if err := rows.Scan(&user.Name, &user.Email, &user.Password); err != nil {
			tr.Rollback()
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		tr.Rollback()
		return nil, err
	}


	return users, tr.Commit()
}

func (u *UserRepo) GetId(id int) (strac.User, error) {
	var user strac.User

	tx, err := u.DB.Begin()
	if err != nil {
		return user, err
	}

	row := tx.QueryRow("SELECT username, email, password FROM users WHERE id = $1", id)
	err = row.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		tx.Rollback()
		return user, err
	}

	return user, tx.Commit()
}
