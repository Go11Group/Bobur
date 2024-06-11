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

	// m := mux.NewRouter()
	m := http.NewServeMux()
	m.HandleFunc("POST /course", h.CreateCourse)
	m.HandleFunc("PUT /course", h.UpdateCourse)
	m.HandleFunc("GET /course/{id}", h.Get)
	m.HandleFunc("GET /courses", h.GetAll)
	m.HandleFunc("DELETE /course/", h.DeleteCourse)

	m.HandleFunc("POST /userCreate", h.CreateUser1)
	m.HandleFunc("PUT /userUpdate", h.UpdateUser1)
	m.HandleFunc("GET /user/", h.GetUser1Id)
	m.HandleFunc("GET /userGetAll", h.GetAllUser1)
	m.HandleFunc("DELETE /user/{id}", h.DeleteCourse1)

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		panic(err)
	}

}
