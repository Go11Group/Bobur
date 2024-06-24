package handler

import (
	"net/http"
	"project/model"

	"github.com/gin-gonic/gin"
)


func (h *Handler) CreateC(g *gin.Context) {
	card := model.Card{}

	err := g.BindJSON(&card)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(card.Number) != 16 {
		g.JSON(http.StatusBadRequest, gin.H{"error":"karta raqami mezon doirasida emas!"})
		return
	}

	err = h.Card.CreateCard(&card)
	if err != nil {
		g.JSON(http.StatusCreated, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, "SUCCESS")
}

func (h *Handler) DeleteC(g *gin.Context) {
	id := g.Param("id")

	err := h.Card.DeleteCard(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

func (h *Handler) UpdateC(g *gin.Context) {
	id := g.Param("id")

	card := model.Card{}

	err := g.BindJSON(&card)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	res, err := h.Card.UpdateCard(id, &card)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	g.JSON(http.StatusOK, res)
}

func (h *Handler) GetCID(g *gin.Context) {
	id := g.Param("id")

	res, err := h.Card.GetCardId(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}

func (h *Handler) GetCAll(g *gin.Context) {
	res, err := h.Card.GetCardAll()

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, res)
}