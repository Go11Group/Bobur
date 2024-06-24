package handler

import (
	"database/sql"
	"model/postgres"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db         *sql.DB
	User       *postgres.UserRepo
	Course     *postgres.CourseRepo
	Lesson     *postgres.LessonRepo
	Enrollment *postgres.EnrollmentRepo
}

func NewHandler(db *sql.DB, user *postgres.UserRepo, course *postgres.CourseRepo, lesson *postgres.LessonRepo, enrollment *postgres.EnrollmentRepo) *Handler {
	return &Handler{db, user, course, lesson, enrollment}
}

func (h *Handler) NewServer() *gin.Engine {


	router := gin.Default()
	// APIS

	// course api
	router.POST("/course", h.CourseCreate)
	router.PUT("/course/:id", h.CourseUpdate)
	router.GET("/course/:id", h.CourseGetId)
	router.GET("/courses", h.CourseGetAllFilter)
	router.DELETE("/course/:id", h.CourseDelete)
	// qo`shimcha api course
	router.GET("/courses/:course_id/lessons/:id", h.CourseGtLessons)
	router.GET("/courses/:course_id/enrollments/:id", h.CourseGetByUsers)
	router.GET("/users/:user_id/courses", h.PopularCourse)

	// user api
	router.POST("/user", h.UserCreate)
	router.PUT("/user/:id", h.UserUpdate)
	router.GET("/user/:id", h.UserGetId)
	router.GET("/users", h.UserGetAllFilter)
	router.DELETE("/user/:id", h.UserDelete)
	// qo`shimcha api user
	router.GET("/users/:user_id/courses/:id", h.UserCoursesGet)
	router.GET("/users/search", h.UsersSearch)

	// lesson api
	router.POST("/lesson", h.LessonCreate)
	router.PUT("/lesson/:id", h.LessonUpdate)
	router.GET("/lesson/:id", h.LessonGetId)
	router.GET("/lessons", h.LessonGetAllFilter)
	router.DELETE("/lesson/:id", h.LessonDelete)

	// enrollment api
	router.POST("/enrollment", h.EnrollmentsCreate)
	router.PUT("/enrollment/:id", h.EnrollmentUpdate)
	router.GET("/enrollment/:id", h.EnrollmentGetId)
	router.GET("/enrollments", h.EnrollmentGetAllFilter)
	router.DELETE("/enrollment/:id", h.EnrollmentDelete)

	return router

}
