package handler

import (
	"fmt"
	"model/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) EnrollmentsCreate(g *gin.Context) {
	enrollment := model.Enrollments{}

	err := g.BindJSON(&enrollment)
	fmt.Println(enrollment, 111)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Enrollment.CreateEnrollment(&enrollment)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

func (h *Handler) EnrollmentUpdate(g *gin.Context) {
	id := g.Param("id")
	enrollment := model.Enrollments{}	

	err := g.BindJSON(&enrollment)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Enrollment.UpdateEnrollment(&enrollment, id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCES")
}

func (h *Handler) EnrollmentDelete(g *gin.Context) {
	id := g.Param("id")

	err := h.Enrollment.DeleteEnrollment(id)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

func (h *Handler) EnrollmentGetAllFilter(g *gin.Context) {
	filter := model.FilterEnrollments{}
	filter.UserId = g.Query("user_id")
	filter.CourseId = g.Query("course_id")
	filter.EnrollmentDate = g.Query("enrollment_date")
	filter.UpdatedAd = g.Query("update_at")
	filter.CreatedAt = g.Query("create_id")	

	enrollment, err := h.Enrollment.GetAllFilterEnrollments(&filter)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, enrollment)

}



func (h *Handler) EnrollmentGetId(g *gin.Context) {
	id := g.Param("id")
	fmt.Println(id)
	enrollment, err := h.Enrollment.GetIdEnrollment(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, enrollment)
}
