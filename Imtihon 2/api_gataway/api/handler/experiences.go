package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"api_gateway/genproto"
)
// CreateExperience
// @Summary Create Experience
// @Description This API for creating experience
// @Tags experience
// @Accept json
// @Produce json
// @Param experience body genproto.CreateExperienceRequest true "Experience"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /experience/create [post]
func (h *Handler) CreateExperience(c *gin.Context) {
	var req genproto.CreateExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := h.Experience.CreateExperience(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Experience created successfully",
	})
}

// UpdateExperience
// @Summary Update Experience
// @Description This API for updating experience
// @Tags experience
// @Accept json
// @Produce json
// @Param experience body genproto.UpdateExperienceRequest true "Experience"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /experience/update [put]
func (h *Handler) UpdateExperience(c *gin.Context) {
	var req genproto.UpdateExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := h.Experience.UpdateExperience(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Experience updated successfully",
	})
}

// DeleteExperience
// @Summary Delete Experience
// @Description This API for deleting experience
// @Tags experience
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /experience/delete/{id} [delete]
func (h *Handler) DeleteExperience(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	_, err := h.Experience.DeleteExperience(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Experience deleted successfully",
	})
}

// GetExperience
// @Summary Get Experience
// @Description This API for getting experience
// @Tags experience
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /experience/get/{id} [get]
func (h *Handler) GetExperience(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	res, err := h.Experience.GetExperience(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Experiences
// @Description This API for getting experiences
// @Tags experience
// @Accept json
// @Produce json
// @Param user_id query string false "Experience User ID"
// @Param company query string false "Company Name"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /experience/getall [get]
func (h *Handler) GetExperiences(c *gin.Context) {
    var req genproto.ExperienceFilter
    req.UserId = c.Query("user_id")
    req.Company = c.Query("company")

    res, err := h.Experience.GetExperiences(context.Background(), &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}

// @Summary Get Experiences By User
// @Description This API for getting experiences
// @Tags experience
// @Accept json
// @Produce json
// @Param user_id query string false "Experience User ID"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /experience/getallbyuser [get]
func (h *Handler) GetExperiencesByUser(c *gin.Context) {
    var req genproto.ExperienceFilter
    req.UserId = c.Query("user_id")

    res, err := h.Experience.GetExperiences(context.Background(), &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}