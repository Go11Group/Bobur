package postgres

import (
	"database/sql"
	"fmt"
	filters "model/filter"
	"model/model"
	"time"
)

type CourseRepo struct {
	Db *sql.DB
}

func NewCourseRwpo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db}
}

// Course yaratish
func (c *CourseRepo) CreateCourse(course *model.Courses) error {
	_, err := c.Db.Exec("insert into courses(title, description)"+
		"values($1, $2)", course.Title, course.Description)

	if err != nil {

		fmt.Printf("Create Query error: %v", err)
		return nil
	}

	return nil
}

// Courseni id orqali ma`lumot ni chiqarish
func (c *CourseRepo) GetCourseId(id string) (model.Courses, error) {
	course := model.Courses{}
	fmt.Println(id)
	row := c.Db.QueryRow("SELECT course_id, title, description, created_at, updated_at FROM courses WHERE course_id=$1", id)
	err := row.Scan(&course.Id, &course.Title, &course.Description, &course.CreateAt, &course.UpdateAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Courses{}, nil
		}
		fmt.Println("Error retrieving course:", err)
		return model.Courses{}, err
	}
	return course, nil

}

// Course ni id orqali Update qilish
func (c *CourseRepo) UpdateCourse(course *model.Courses, id string) (model.Courses, error) {

	_, err := c.Db.Exec("update courses set title=$1, description=$2, updated_at=$4  where course_id=$3", course.Title, course.Description, id, time.Now())
	if err != nil {
		return model.Courses{}, nil
	}
	res, err := c.GetCourseId(id)
	if err != nil {
		return model.Courses{}, nil
	}
	return res, nil
}

// Course ni id orqali delete qilish
func (c *CourseRepo) DeleteCourse(id string) error {
	_, err := c.Db.Exec("delete from courses where course_id=$1", id)
	if err != nil {
		return nil
	}
	return nil
}

// Course ni ma`lum bir columinlari orqali filterlash
func (c *CourseRepo) GetAllCourseFilter(f *model.FilterCourse) ([]model.Courses, error) {
	filter := ""
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)

	query := `select course_id, title, description, created_at
	 	from courses where true `

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title "
	}

	if len(f.Description) > 0 {
		params["description"] = f.Description
		filter += " and description = :description "
	}

	if len(f.Limit) > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	query = query + filter + limit

	query, arr = filters.ReplaceQueryParams(query, params)
	rows, err := c.Db.Query(query, arr...)
	if err != nil {
		panic(err)
	}

	courses := []model.Courses{}
	for rows.Next() {
		course := model.Courses{}
		err = rows.Scan(&course.Id, &course.Title, &course.Description, &course.CreateAt)
		if err != nil {
			fmt.Println("2222 filtersda xato bor course")
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

// Ma'lum bir kursga tegishli barcha darslarni olish.
func (u *CourseRepo) GetLessonsbyCourse(id string) (model.GetLessonsbyCourse, error) {
	courseL := model.GetLessonsbyCourse{}
	query := `select l.lesson_id, l.title, l.content from lessons as l 
	left join courses as c on c.course_id=l.course_id
	where l.course_id=` + "'" + id + "'"

	rows, err := u.Db.Query(query)
	if err != nil {
		return courseL, err
	}
	courseL.Course_id = id
	for rows.Next() {
		lesson := model.Lessons1{}
		err = rows.Scan(&lesson.Id, &lesson.Title, &lesson.Content)
		if err != nil {
			return model.GetLessonsbyCourse{}, err
		}
		courseL.Lessons = append(courseL.Lessons, lesson)
	}

	return courseL, nil

}

// Ma'lum bir kursga ro'yxatdan o'tgan barcha foydalanuvchilarni olish.
func (c *CourseRepo) GetEnrolledUsersbyCourse(id string) (model.GetEnrolledUsersbyCourse, error) {
	courseUE := model.GetEnrolledUsersbyCourse{}

	query := `select u.user_id, u.name, u.email from enrollments as e 
	inner join users as u on u.user_id=e.user_id 
	where e.course_id=` + "'" + id + "'"

	rows, err := c.Db.Query(query)
	if err != nil {
		return courseUE, err
	}

	courseUE.CourseId = id
	for rows.Next() {
		user := model.User1{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return model.GetEnrolledUsersbyCourse{}, err
		}
		courseUE.User1 = append(courseUE.User1, user)
	}

	return courseUE, nil
}


// Ma'lum bir kursga ro'yxatdan o'tgan barcha foydalanuvchilarni olish.
func (repo *CourseRepo) GetPopularyCourse(startTime, endTime time.Time) ([]map[string]interface{}, error) {

	query := `select c.course_id, c.title, COUNT(e.enrollment_id) as enrollments_count from courses c 
	join enrollments e on c.course_id = e.course_id where e.enrollment_date 
	between $1 and $2 group by  c.course_id, c.title order by  enrollments_count desc `
	rows, err := repo.Db.Query(query, startTime, endTime)
	if err != nil {
		return nil, err
	}
	var popularCourses []map[string]interface{}
	var courseID, courseTitle, enrollmentsCount string
	for rows.Next() {
		if err := rows.Scan(&courseID, &courseTitle, &enrollmentsCount); err != nil {
			return nil, err
		}
		popularCourses = append(popularCourses, map[string]interface{}{
			"course_id":         courseID,
			"course_title":      courseTitle,
			"enrollments_count": enrollmentsCount,
		})
	}
	return popularCourses, nil

}
