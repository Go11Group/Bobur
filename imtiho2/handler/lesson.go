package handler

import (
	"fmt"
	"model/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Posmanda ma'lumot olib sql bn bo`g`langan papkaga yuboradi
func (h *Handler) LessonCreate(g *gin.Context) {
	lesson := model.Lessons{}

	err := g.BindJSON(&lesson)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Lesson.CreateLesson(&lesson)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

// Posmandan ma'lumot oladi va sqlga yuboradi
func (h *Handler) LessonUpdate(g *gin.Context) {
	id := g.Param("id")
	fmt.Println(id)
	lesson := model.Lessons{}
	
	err := g.BindJSON(&lesson)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(1)
	err = h.Lesson.UpdateLesson(&lesson, id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCES")
}

// posmandan id olib sqlga yuboradi va uni o'chiradi
func (h *Handler) LessonDelete(g *gin.Context) {
	id := g.Param("id")
	err := h.Lesson.DeleteLesson(id)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "SUCCESS")
}

// posman paramdan ma'lumot olib sqlga berib yuboradi va kegan ma'lumot larni tekshiradi
func (h *Handler) LessonGetAllFilter(g *gin.Context) {
	filter := model.FilterLesson{}
	filter.CourseId = g.Query("course_id")
	filter.Title = g.Query("title")
	filter.Content = g.Query("content")
	fmt.Println(filter.Content, filter.Title)

	lesson, err := h.Lesson.GetAllFilterLesson(&filter)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, lesson)

}

// Posmanda id keladi va usha id orqali ma'lumotlarni 
func (h *Handler) LessonGetId(g *gin.Context) {
	id := g.Param("id")

	lesson, err := h.Lesson.GetIdLesson(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, lesson)
}
