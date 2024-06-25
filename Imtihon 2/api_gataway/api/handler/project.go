package handler

import (
	"context"
	"net/http"

	"api_gateway/genproto"

	"github.com/gin-gonic/gin"
)

// CreateProject
// @Summary Create Project
// @Description This API for creating project
// @Tags project
// @Accept json
// @Produce json
// @Param project body genproto.CreateProjectRequest true "Project"
// @Success 201 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /project/create [post]
func (h *Handler) CreateProject(c *gin.Context) {
	var req genproto.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := h.Project.CreateProject(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
	})
}

// UpdateProject
// @Summary Update Project
// @Description This API for updating project
// @Tags project
// @Accept json
// @Produce json
// @Param project body genproto.UpdateProjectRequest true "Project"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /project/update [put]
func (h *Handler) UpdateProject(c *gin.Context) {
	var req genproto.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := h.Project.UpdateProject(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Project updated successfully",
	})
}

// DeleteProject
// @Summary Delete Project
// @Description This API for deleting project
// @Tags project
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /project/delete/{id} [delete]
func (h *Handler) DeleteProject(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	_, err := h.Project.DeleteProject(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Project deleted successfully",
	})
}

// GetProject
// @Summary Get Project
// @Description This API for getting project
// @Tags project
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /project/get/{id} [get]
func (h *Handler) GetProject(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	res, err := h.Project.GetProject(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Projects
// @Description This API for getting projects
// @Tags project
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Param url query string false "Project url"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /project/getall [get]
func (h *Handler) GetProjects(c *gin.Context) {
	var req genproto.ProjectFilter
	req.UserId = c.Query("user_id")
	req.Url = c.Query("url")

	res, err := h.Project.GetProjects(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}


// @Summary Get Projects By User
// @Description This API for getting projects
// @Tags project
// @Accept json
// @Produce json
// @Param user_id query string false "User ID"
// @Success 200 {object} string "ok"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /project/getallbyuser [get]
func (h *Handler) GetProjectsByUser(c *gin.Context) {
	var req genproto.ProjectFilter
	req.UserId = c.Query("user_id")

	res, err := h.Project.GetProjects(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
