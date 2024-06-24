package model

type Lessons struct {
	Id        string `json:"id"`
	CourseId  string `json:"course_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_id"`
}
