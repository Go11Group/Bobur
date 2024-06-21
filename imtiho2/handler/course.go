package handler

import (
	"fmt"
	"model/model"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Posmanda mulumotni olib sql ga berib yuboradi
func (h *Handler) CourseCreate(g *gin.Context) {
	course := model.Courses{}

	err := g.BindJSON(&course)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Course.CreateCourse(&course)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

// posmandan mulumot kelib sqlga burish func
func (h *Handler) CourseUpdate(g *gin.Context) {
	id := g.Param("id")
	fmt.Println(id)
	course := model.Courses{}

	err := g.BindJSON(&course)
	if err != nil {
		g.IndentedJSON(http.StatusAccepted, err)
		return
	}

	res, err := h.Course.UpdateCourse(&course, id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, res)
}

// posmandan id kelai va sql ga beriladi va shu id orqali delete qiladi
func (h *Handler) CourseDelete(g *gin.Context) {
	id := g.Param("id")

	err := h.Course.DeleteCourse(id)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

// pomandan ma`lum bir malumotlar keladi va sql ga yuboradi
func (h *Handler) CourseGetAllFilter(g *gin.Context) {
	filter := model.FilterCourse{}
	filter.Title = g.Query("title")
	filter.Description = g.Query("description")
	filter.Limit = g.Query("limit")
	fmt.Println(filter.Title, filter.Description)

	courses, err := h.Course.GetAllCourseFilter(&filter)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, courses)

}

func (h *Handler) CourseGtLessons(g *gin.Context) {
	id := g.Param("id")

	courseL, err := h.Course.GetLessonsbyCourse(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, courseL)
}

func (h *Handler) CourseGetByUsers(g *gin.Context) {
	id := g.Param("id")

	courseUE, err := h.Course.GetEnrolledUsersbyCourse(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, courseUE)

}

// posmandagi id ni oladi va sql ga yuboradi
func (h *Handler) CourseGetId(g *gin.Context) {
	id := g.Param("id")
	fmt.Println(id)
	course, err := h.Course.GetCourseId(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, course)
}

// Ma'lum bir muddat davomida eng ko'p ro'yxatdan o'tilgan kurslarni olish
func (h *Handler) PopularCourse(g *gin.Context) {

	t1 := g.Query("start_time")
	t2 := g.Query("end_time")

	startTime1 := strings.Split(t1, "-")

	startYear, err := strconv.Atoi(startTime1[0])
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startMonth, err := strconv.Atoi(startTime1[1])
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	startDay, err := strconv.Atoi(startTime1[2])
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endTime1 := strings.Split(t2, "-")

	endYear, err := strconv.Atoi(endTime1[0])
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endMonth, err := strconv.Atoi(endTime1[1])
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endDay, err := strconv.Atoi(endTime1[2])
	if err != nil {
		panic(err)
	}

	startTime := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(endYear, time.Month(endMonth), endDay, 0, 0,
		0, 0, time.UTC)

	courses, err := h.Course.GetPopularyCourse(startTime, endTime)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"time_period": gin.H{
			"start_date": startTime,
			"end_date":   endTime,
		},
		"popular_courses": courses,
	}
	g.JSON(http.StatusOK, response)
}
