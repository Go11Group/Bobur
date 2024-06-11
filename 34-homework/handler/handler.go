package handler

import (
	"database/sql"
	"fmt"
	"model/postgres"
	"net/http"
)

type Handler struct {
	db     *sql.DB
	Course *postgres.CourseRepo
	User   *postgres.UserRepo
}

func NewHandler(db *sql.DB, course *postgres.CourseRepo, user *postgres.UserRepo) *Handler {
	return &Handler{db, course, user}
}

func (h *Handler) NewServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /course", h.CreateCourse)
	mux.HandleFunc("PUT /course", h.UpdateCourse)
	mux.HandleFunc("GET /course/{id}", h.Get)
	mux.HandleFunc("GET /courses", h.GetAll)
	mux.HandleFunc("DELETE /course/", h.DeleteCourse)

	mux.HandleFunc("POST /user", h.CreateUser1)
	mux.HandleFunc("PUT /user", h.UpdateUser1)
	mux.HandleFunc("GET /user/", h.GetUser1Id)
	mux.HandleFunc("GET /users", h.GetAllUser1)
	mux.HandleFunc("DELETE /user/", h.DeleteCourse1)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}
