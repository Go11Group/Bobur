package model

type FilterUser struct {
	Limit    string `json:"limit"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type FilterCourse struct {
	Limit       string `json:"limit"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Create_at   string `json:"create_at"`
	Update_at   string `json:"update_at"`
}

type FilterLesson struct {
	Limit     string `json:"limit"`
	CourseId  string `json:"course_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_id"`
}

type FilterEnrollments struct {
	Limit           string `json:"limit"`
	UserId          string `json:"user_id"`
	CourseId        string `json:"course_id"`
	EnrollmentDate string `json:" enrollment_date"`
	CreatedAt       string `json:"created_at"`
	UpdatedAd       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}
