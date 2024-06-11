package handler

import (
	"encoding/json"
	"model/models"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) CreateUser1(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.User.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateUser1(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.User.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeleteCourse1(w http.ResponseWriter, r *http.Request) {
	d := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(d)
	if err != nil {
		panic(err)
	}

	err = h.User.DeleteUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetAllUser1(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}

	users, err := h.User.GetAllUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetUser1Id(w http.ResponseWriter, r *http.Request) {
	d := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(d)
	if err != nil {
		panic(err)
	}
	user := models.User{}

	user, err = h.User.GetUserId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
