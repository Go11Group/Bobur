package postgres

import (
	"database/sql"
	"fmt"
	filters "model/filter"
	"model/model"
	"time"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// User yaratish 
func (u *UserRepo) CreateUser(user *model.Users) error {
	parsedTime, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		fmt.Println("Vaqtni o'zgartirishda xatolik:", err)
		return nil
	}
	_, err = u.Db.Exec("INSERT INTO users(name, email, birthday, password, created_at)"+
		"VALUES($1, $2, $3, $4, $5)", user.Name, user.Email, parsedTime, user.Password, time.Now())

	if err != nil {
		fmt.Printf("Create Query error: %v", err)
		return err
	}

	return nil
}

// ma`lum bir id bo`yicha qidirish
func (u *UserRepo) GetUserId(id string) (model.Users, error) {
	fmt.Println(id)
	user := model.Users{}
	row := u.Db.QueryRow("select user_id, name, email, birthday, password, created_at, updated_at from users where user_id=$1", id)

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		fmt.Println("User id si orqali")
		return model.Users{}, err
	}
	fmt.Println(user)
	return user, nil
}

//  Ma`lum bir id bo`yicha yangilash 
func (u *UserRepo) UpdateUser(user *model.Users, id string) error {
	_, err := u.Db.Exec("update users set name=$1, email=$2, birthday=$3, password=$4, updated_at where id=$5", user.Name, user.Email, user.Birthday, user.Password, id)
	if err != nil {
		return nil
	}
	return nil
}

// Ma`lum bir id bo`yicha o`chirish
func (u *UserRepo) DeleteUser(id string) error {
	now := time.Now()
	unixTime := now.Unix()
	DeletedAt := int(unixTime)
	_, err := u.Db.Exec("DELETE FROM users WHERE user_id=$1", id)
	if err != nil {
		return err
	}

	_, err = u.Db.Exec("UPDATE users SET deleted_at=$1 WHERE user_id=$2", DeletedAt, id)
	if err != nil {
		return nil
	}

	return nil
}

// columinlar bo`yicha filterlash 
func (u *UserRepo) GetAllUserFilter(f *model.FilterUser) ([]model.Users, error) {
	filter := ""
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit string
	)

	query := `select user_id, name, email, birthday, password, created_at, updated_at, deleted_at
	 	from users where true`

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += " and name = :name "
	}

	if len(f.Email) > 0 {
		params["email"] = f.Email
		filter += " and email = :email "
	}

	if len(f.Birthday) > 0 {
		params["birthday"] = f.Birthday
		filter += " and birthday = :birthday "
	}

	if len(f.CreateAt) > 0 {
		params["created_at"] = f.CreateAt
		filter += " and created_at = :created_at "
	}

	if len(f.UpdateAt) > 0 {
		params["updated_at"] = f.UpdateAt
		filter += " and updated_at = :updated_at "
	}

	if len(f.Limit) > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	query = query + filter + limit

	query, arr = filters.ReplaceQueryParams(query, params)
	rows, err := u.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	users := []model.Users{}
	for rows.Next() {
		user := model.Users{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreateAt, &user.UpdateAt, &user.DeleteAt)
		if err != nil {
			fmt.Println("rowsda xato")
			return nil, err
		}
		users = append(users, user)
	}
	fmt.Println(users)
	return users, nil
}

// Ma'lum bir foydalanuvchiga tegishli barcha kurslarni olish
func (u *UserRepo) GetCoursesbyUser(id string) (model.GetCoursesByUser, error) {

	usersC := model.GetCoursesByUser{}
	query := `SELECT c.course_id, c.title, c.description
	FROM enrollments AS e
	INNER JOIN courses AS c ON c.course_id = e.course_id
	WHERE e.user_id=` + "'" + id + "'"

	rows, err := u.Db.Query(query)
	if err != nil {
		return usersC, err
	}
	usersC.UserId = id
	for rows.Next() {
		course := model.Courses1{}
		err = rows.Scan(&course.Id, &course.Title, &course.Description)
		if err != nil {
			return model.GetCoursesByUser{}, err
		}
		usersC.Courses = append(usersC.Courses, course)
	}

	return usersC, nil
}

// name yoki email orqali qidirish 
func (u *UserRepo) SearchUsers(f model.FilterUser) ([]model.User1, error) {
	filter := ""
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit string
	)

	query := `select u.user_id, u.name, u.email from users as u 
	where true `

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += " and name = :name "
	}

	if len(f.Email) > 0 {
		params["email"] = f.Email
		filter += " and email = :email "
	}

	if len(f.Limit) > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}
	query = query + filter + limit
	fmt.Println(query)
	query, arr = filters.ReplaceQueryParams(query, params)
	rows, err := u.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	users := []model.User1{}
	for rows.Next() {
		user := model.User1{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
