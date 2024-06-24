package postgres

import (
	"database/sql"
	"fmt"
	filters "model/filter"
	"model/model"
	"time"
)

type LessonRepo struct {
	Db *sql.DB
}

func NewLessonRwpo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db}
}

// Lesson yaratish
func (l *LessonRepo) CreateLesson(lesson *model.Lessons) error {
	_, err := l.Db.Exec("INSERT INTO lessons(course_id, title, content, created_at)"+
		"VALUES($1, $2, $3, $4)", lesson.CourseId, lesson.Title, lesson.Content, time.Now())

	if err != nil {
		return err
	}
	return nil
}


// Lesson ni Update qilish
func (l *LessonRepo) UpdateLesson(lesson *model.Lessons, id string) error {
	fmt.Println(2)
	_, err := l.Db.Exec("update lessons set course_id=$1, title=$2, updated_at=$3 where lesson_id=$4", lesson.CourseId, lesson.Title, time.Now(), id)
	if err != nil {
		fmt.Printf("Update Query error: %v", err)
		return err
	}

	return nil
}

// Lessonni Delete qilish id orqali
func (l *LessonRepo) DeleteLesson(id string) error {
	now := time.Now()
	unixTime := now.Unix()
	DeletedAt := int(unixTime)
	fmt.Println(DeletedAt)
	_, err := l.Db.Exec("delete from lessons where lesson_id=$1", id)
	if err != nil {
		fmt.Printf("Create Query error: %v", err)
		panic(err)
	}

	return nil
}


// Columinlar orqali filterlash Lesson tableni
func (l *LessonRepo) GetAllFilterLesson(f *model.FilterLesson) ([]model.Lessons, error) {
	filter := ""
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)

	query := `select lesson_id, course_id, title, content, created_at, updated_at, deleted_at
	 	from lessons where true `

	if len(f.Limit) > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title "
	}

	if len(f.CourseId) > 0 {
		params["course_id"] = f.CourseId
		filter += " and course_id = :course_id "
	}

	if len(f.Content) > 0 {
		params["content"] = f.Content
		filter += " and content = :content "
	}

	if len(f.CreatedAt) > 0 {
		params["created_at"] = f.CreatedAt
		filter += " and created_at = :created_at "
	}

	if len(f.UpdatedAt) > 0 {
		params["updated_at"] = f.UpdatedAt
		filter += " and updated_at = :updated_at "
	}

	query = query + filter + limit

	query, arr = filters.ReplaceQueryParams(query, params)
	fmt.Println(query)

	rows, err := l.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	lessons := []model.Lessons{}
	for rows.Next() {
		lesson := model.Lessons{}
		err = rows.Scan(&lesson.Id, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)

		if err != nil {
			fmt.Println("rowsda xato", err)
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	fmt.Println(lessons)
	return lessons, nil
}

// Lessonni id orqali bitta ma`lumotini chiqarish 
func (l *LessonRepo) GetIdLesson(id string) (model.Lessons, error) {
	lesson := model.Lessons{}
	fmt.Println(id, 11)
	row := l.Db.QueryRow("select lesson_id, course_id, title, content, created_at, updated_at from lessons where lesson_id=$1", id)
	err := row.Scan(&lesson.Id, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt)
	if err != nil {
		return model.Lessons{}, err
	}
	return lesson, nil

}
