package model

type GetCoursesByUser struct {
	UserId string     `json:"user_id"`
	Courses []Courses1 `json:"courses"`
}

type GetLessonsbyCourse struct {
	Course_id string     `json:"course_id"`
	Lessons   []Lessons1 `json:"lessons"`
}

type Lessons1 struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Courses1 struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type User1 struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetEnrolledUsersbyCourse struct {
	CourseId string  `json:"course_id"`
	User1     []User1 `json:"enrolled_users"`
}
