package handler

import (
	"net/http"
	"project/model"

	"github.com/gin-gonic/gin"
)


func (h *Handler) CreateTR(g *gin.Context) {
	tr := model.Transaction{}

	err := g.BindJSON(&tr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Transaction.CreateTR(&tr)
	if err != nil {
		g.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, "SUCCESS")
}

func (h *Handler) DeleteTR(g *gin.Context) {
	id := g.Param("id")

	err := h.Transaction.DeleteTR(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

func (h *Handler) UpdateTR(g *gin.Context) {
	id := g.Param("id")

	tr := model.Transaction{}

	err := g.BindJSON(&tr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	res, err := h.Transaction.UpdateTR(id, &tr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	g.JSON(http.StatusOK, res)
}

func (h *Handler) GetTRID(g *gin.Context) {
	id := g.Param("id")

	res, err := h.Transaction.GetTRByID(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}

func (h *Handler) GetTRAll(g *gin.Context) {
	res, err := h.Transaction.GetTRAll()

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, res)
}