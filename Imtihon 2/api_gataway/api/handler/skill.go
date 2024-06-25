package handler

import (
	"context"
	"net/http"

	"api_gateway/genproto"

	"github.com/gin-gonic/gin"
)

// CreateSkill
// @Summary Create Skill
// @Description This API for creating skill
// @Tags skill
// @Accept json
// @Produce json
// @Param skill body genproto.CreateSkillRequest true "Skill"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /skill/create [post]
func (h *Handler) CreateSkill(c *gin.Context) {
	var req genproto.CreateSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := h.Skill.CreateSkill(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Skill created successfully",
	})
}

// UpdateSkill
// @Summary Update Skill
// @Description This API for updating skill
// @Tags skill
// @Accept json
// @Produce json
// @Param skill body genproto.UpdateSkillRequest true "Skill"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /skill/update [put]
func (h *Handler) UpdateSkill(c *gin.Context) {
	var req genproto.UpdateSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := h.Skill.UpdateSkill(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Skill updated successfully",
	})
}

// DeleteSkill
// @Summary Delete Skill
// @Description This API for deleting skill
// @Tags skill
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /skill/delete/{id} [delete]
func (h *Handler) DeleteSkill(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	_, err := h.Skill.DeleteSkill(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Skill deleted successfully",
	})
}

// GetSkill
// @Summary Get Skill
// @Description This API for getting skill
// @Tags skill
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /skill/get/{id} [get]
func (h *Handler) GetSkill(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	res, err := h.Skill.GetSkill(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Skills
// @Description This API for getting skills
// @Tags skill
// @Accept json
// @Produce json
// @Param name query string false "Skill Name"
// @Param user_id query string false "Skill Level"
// @Success 200 {object} genproto.Skills
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /skill/getall [get]
func (h *Handler) GetSkills(c *gin.Context) {
	var req genproto.SkillFilter
	req.Name = c.Query("name")
	req.UserId = c.Query("user_id")

	res, err := h.Skill.GetSkills(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}


// @Summary Get Skills By User
// @Description This API for getting skills
// @Tags skill
// @Accept json
// @Produce json
// @Param user_id query string false "Skill User Id"
// @Success 200 {object} genproto.Skills
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /skill/getallbyuser [get]
func (h *Handler) GetSkillsByUser(c *gin.Context) {
	var req genproto.SkillFilter
	req.UserId = c.Query("user_id")

	res, err := h.Skill.GetSkills(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
