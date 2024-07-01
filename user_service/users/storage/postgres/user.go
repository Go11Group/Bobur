package postgres

import (
	"database/sql"
	"time"

	"user_service/models"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepositoryRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo UserRepository) CreateUser(user *models.Users) (interface{}, error) {
	return repo.Db.Exec("insert into users(username,email,password_hash,created_at)",user.UserName, user.Email, user.Password_hash, time.Now())
}
func (repo UserRepository) UpdateUser(user *models.Users, id string) (interface{}, error) {
	return repo.Db.Exec("update users set username=$1,email=$2,password_hash=$3,updated_at=$4 where id=$5 and deleted_at=0)", user.UserName, user.Email, user.Password_hash, time.Now(), id)
}
func (repo UserRepository) DeleteUser(id string) (interface{}, error) {
	return repo.Db.Exec("update users set deleted_at=$1 where id=$2 and deleted_at is null)", time.Now(), id)
}
func (repo UserRepository) GetUserById(id string) (*models.Users, error) {
	rows, err := repo.Db.Query("select username,email ,password_hash from users  where id=$1 and deleted_at is null)", id)
	if err != nil {
		return nil, err
	}
	var user models.Users
	for rows.Next() {
		err := rows.Scan(&user.UserName, &user.Email, &user.Password_hash)
		if err != nil {
			return nil, err
		}
		// return nil, nil
	}
	return nil, err
}
func (repo UserRepository) GetUser(user *models.Users) (*[]models.Users, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(user.UserName) > 0 {
		params["username"] = user.UserName
		filter += " and username = :username "

	}
	if len(user.Email) > 0 {
		params["email"] = user.Email
		filter += " and email = :email "

	}
	if len(user.Password_hash) > 0 {
		params["password_hash"] = user.Password_hash
		filter += " and password_hash = :password_hash "

	}

	if user.Limit > 0 {
		params["limit"] = user.Limit
		limit = ` LIMIT :limit`

	}
	if user.Offset > 0 {
		params["offset"] = user.Offset
		limit = ` OFFSET :offset`

	}
	query := "select id,username,email,password_hash from users  where  deleted_at is null"

	query = query + filter + limit + offset
	query, arr = ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	var users []models.Users
	for rows.Next() {
		var user models.Users
		err := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password_hash)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, err
}
