package handler

import (
	"api_gateway/genproto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Education
// @Description This API for creating education
// @Tags education
// @Accept json
// @Produce json
// @Param education body genproto.CreateEducationRequest true "Education"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /education/create [post]
func (h *Handler) CreateEducation(c *gin.Context) {
	var req genproto.CreateEducationRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Education.CreateEducation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Education created successfully",
	})
}

// UpdateEducation
// @Summary Update Education
// @Description This API for updating education
// @Tags education
// @Accept json
// @Produce json
// @Param education body genproto.UpdateEducationRequest true "Education"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /education/update [put]
func (h *Handler) UpdateEducation(c *gin.Context) {
	var req genproto.UpdateEducationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := h.Education.UpdateEducation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Education updated successfully",
	})
}

// DeleteEducation
// @Summary Delete Education
// @Description This API for deleting education
// @Tags education
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /education/delete/{id} [delete]
func (h *Handler) DeleteEducation(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	_, err := h.Education.DeleteEducation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Education deleted successfully",
	})
}
// GetEducation
// @Summary Get Education
// @Description This API for getting education
// @Tags education
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} genproto.Education
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /education/get/{id} [get]
func (h *Handler) GetEducation(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	res, err := h.Education.GetEducation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Educations
// @Description This API for getting educations
// @Tags education
// @Accept json
// @Produce json
// @Param user_id query string false "UserId"
// @Param institution query string false "Institution"
// @Param degree query string false "Degree"
// @Success 200 {object} genproto.Education
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /education/getall [get]
func (h *Handler) GetEducations(c *gin.Context) {
    var req genproto.EducationFilter
	req.UserId = c.Query("user_id")
    req.Institution = c.Query("institution")
    req.Degree = c.Query("degree")
    res, err := h.Education.GetEducations(context.Background(), &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}


// @Summary Get Educations By UserId
// @Description This API for getting educations
// @Tags education
// @Accept json
// @Produce json
// @Param user_id query string false "UserId"
// @Success 200 {object} genproto.Education
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /education/getallbyuser [get]
func (h *Handler) GetEducationsByUser(c *gin.Context) {
    var req genproto.EducationFilter
	req.UserId = c.Query("user_id")

    res, err := h.Education.GetEducations(context.Background(), &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, res)
}
