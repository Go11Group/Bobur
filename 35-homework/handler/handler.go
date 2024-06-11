package handler

import (
	"database/sql"
	"fmt"
	"model/postgres"
	"net/http"

	"github.com/gorilla/mux"
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

	m := mux.NewRouter()
	// m := http.NewServeMux()
	m.HandleFunc("/course", h.CreateCourse).Methods("POST")
	m.HandleFunc("/course", h.UpdateCourse).Methods("PUT")
	m.HandleFunc("/course/{id}", h.Get).Methods("GET")
	m.HandleFunc("/courses", h.GetAll).Methods("GET")
	m.HandleFunc("/course/{id}", h.DeleteCourse).Methods("DELETE ")

	m.HandleFunc("/userCreate", h.CreateUser1).Methods("POST")
	m.HandleFunc("/userUpdate", h.UpdateUser1).Methods("PUT")
	m.HandleFunc("/user/{id}", h.GetUser1Id).Methods("GET")
	m.HandleFunc("/userGetAll", h.GetAllUser1).Methods("GET")
	m.HandleFunc("/user/{id}", h.DeleteCourse1).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", m)
	if err != nil {
		panic(err)
	}

}
