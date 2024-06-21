package handler

import (
	"fmt"
	"model/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


// Posman dan ma`lumot o`qib olib uni psqlga beradi va yangi mumot qo`shiladi tablega
func (h *Handler) UserCreate(g *gin.Context) {
	user := model.Users{}	

	err := g.BindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.User.CreateUser(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return	
	}
	g.JSON(http.StatusOK, "SUCCESS")
}

// Posmanda ma'lumot o'qib olinadi psqlga buradi va update qiladi ID orqali
func (h *Handler) UserUpdate(g *gin.Context) {
	id := g.Param("id")
	user := model.Users{}

	err := g.BindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.User.UpdateUser(&user, id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")

}

// Posman da id olib user tablega beradi va o'chiradi yani connection bo`lgan joyga
func (h *Handler) UserDelete(g *gin.Context) {
	id := g.Param("id")
	fmt.Println(id)
	err := h.User.DeleteUser(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")

}

// posmandan belgilangan ma`lumotlar keladi va sql ga o`sh orqali ma'lumot qaytadi
func (h *Handler) UserGetAllFilter(g *gin.Context) {
	filter := model.FilterUser{}
	filter.Limit = g.Query("limit")
	filter.Name = g.Query("name")	
	filter.Email = g.Query("email")
	fmt.Println(filter.Name, filter.Email)
	// filter.Birthday = g.Param("birthday")
	// filter.Create_at = g.Param("create_at")
	// filter.Update_at = g.Param("update_at")

	users, err := h.User.GetAllUserFilter(&filter)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, users)
}

func (h *Handler) UserGetId(g *gin.Context) {
	id := g.Param("id")
	fmt.Println(id)
	user, err := h.User.GetUserId(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, user)
}

func (h *Handler) UserCoursesGet(g *gin.Context) {
	id := g.Param("id")
	userC, err := h.User.GetCoursesbyUser(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, userC)
}

func (h *Handler) UsersSearch(g *gin.Context) {
	filter := model.FilterUser{}
	filter.Name = g.Query("name")
	filter.Email = g.Query("email")
	filter.Limit = g.Query("limit")

	users, err := h.User.SearchUsers(filter)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, users)
}

