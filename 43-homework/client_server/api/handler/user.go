package handler

import (
	"fmt"
	"model/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

// URL orqali kelgan ma`lumotni i`qish
func (h *Handler) Create(g *gin.Context) {
	user := model.User{}
	err := g.BindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
		fmt.Println(1)
		return
	}

	err = h.User.CreateUser(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(2)
		return
	}

	g.JSON(http.StatusCreated, "SUCCESS")
}

// URL orqali kerlgan ma'lumotni id si orqali boshqa bir ma'lumotga o'zgartirish
func (h *Handler) Update(g *gin.Context) {
	id := g.Param("id")

	user := model.User{}
	err := g.BindJSON(&user)
	if err != nil {
		fmt.Println(1)
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.UpdateUser(id, &user)
	if err != nil {
		g.JSON(http.StatusBadGateway, nil)
		return
	}

	g.JSON(http.StatusAccepted, res)
}

// URL da id keladi va shu id orqali ma'lumot o'chiriladi
func (h *Handler) Delete(g *gin.Context) {
	id := g.Param("id")

	err := h.User.DeleteUser(id)
	if err != nil {
		fmt.Println(1)
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

//  URL da id olib o'sh id ga tegishli ma'lumotlarni chiqarish
func (h *Handler) GetById(g *gin.Context) {
	id := g.Param("id")

	user, err := h.User.GetById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, user)
}

// URL orqali hamma ma'lumotlarni olib kelish 
func (h *Handler) GetAll(g *gin.Context) {
	users, err := h.User.GetByAll()
	if err != nil {
		fmt.Println(1)
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, users)
}

