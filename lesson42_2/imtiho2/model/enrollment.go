package model

type Enrollments struct {
	Id              string `json:"id"`
	UserId         string `json:"user_id"`
	CourseId       string `json:"course_id"`
	EnrollmentDate string `json:"enrollment_date"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}
