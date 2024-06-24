package handler

import (
	"fmt"
	"net/http"
	"project/model"

	"github.com/gin-gonic/gin"
)

// Posmandan ma'lumot oladi va postgresga jonatadi
func (h *Handler) CreateST(g *gin.Context) {
	st := model.Station{}

	err := g.BindJSON(&st)
	fmt.Println(st)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	
	fmt.Println("salom")
	err = h.Station.STCreate(&st)
	if err != nil {
		g.JSON(http.StatusCreated, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, "SUCCESS")
}

// Paramdan id olib o'sha id orqali ma'lumot o'chiradi
func (h *Handler) DeleteST(g *gin.Context) {
	id := g.Param("id")

	err := h.Station.STDelete(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

// Ma'lumot keladi json formatta va shu ma'lumotni postgresga beradi
func (h *Handler) UpdateST(g *gin.Context) {
	id := g.Param("id")

	tr := model.Station{}

	err := g.BindJSON(&tr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Station.STUpdate(id, &tr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}

// Paramdan id oladi va shu id orqali ma'lumotni olib keladi
func (h *Handler) GetSTID(g *gin.Context) {
	id := g.Param("id")

	res, err := h.Station.STGetId(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}

// Postgresdagi ma'lum bir tableni hamma ma'lumotini olib keladi
func (h *Handler) GetSTAll(g *gin.Context) {
	res, err := h.Station.STGetAll()

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, res)
}
