package handler

import (
	"net/http"
	"project/model"

	"github.com/gin-gonic/gin"
)

// Posmandan ma'lumot oladi va postgresga jonatadi
func (h *Handler) CreateT(g *gin.Context) {
	tr := model.Terminal{}

	err := g.BindJSON(&tr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Terminal.CreateT(&tr)
	if err != nil {
		g.JSON(http.StatusCreated, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, "SUCCESS")
}

// Paramdan id olib o'sha id orqali ma'lumot o'chiradi
func (h *Handler) DeleteT(g *gin.Context) {
	id := g.Param("id")

	err := h.Terminal.DeleteT(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

// Ma'lumot keladi json formatta va shu ma'lumotni postgresga beradi
func (h *Handler) UpdateT(g *gin.Context) {
	id := g.Param("id")

	tr := model.Terminal{}

	err := g.BindJSON(&tr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	res, err := h.Terminal.UpdateT(id, &tr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	g.JSON(http.StatusOK, res)
}

// Paramdan id oladi va shu id orqali ma'lumotni olib keladi
func (h *Handler) GetTID(g *gin.Context) {
	id := g.Param("id")

	res, err := h.Terminal.GetTId(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}

// Postgresdagi ma'lum bir tableni hamma ma'lumotini olib keladi
func (h *Handler) GetTAll(g *gin.Context) {
	res, err := h.Terminal.GetTAll()

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, res)
}