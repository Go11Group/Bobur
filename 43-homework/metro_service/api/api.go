package api

import (
	"database/sql"
	"project/api/handler"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	router := gin.Default()

	h := handler.NewHandler(db)
	// card api
	router.POST("card/create", h.CreateC)
	router.PUT("card/update/:id", h.UpdateC)
	router.GET("card/getbyID/:id", h.GetCID)
	router.GET("card/getAll", h.GetCAll)
	router.DELETE("card/delete/:id", h.DeleteC)

	// transaction api
	router.POST("tr/create", h.CreateTR)
	router.PUT("tr/update/:id", h.UpdateTR)
	router.GET("tr/getbyID/:id", h.GetTRID)
	router.GET("tr/getAll", h.GetTRAll)
	router.DELETE("tr/delete/:id", h.DeleteTR)

	// Terminal api
	router.POST("t/create", h.CreateT)
	router.PUT("t/update/:id", h.UpdateT)
	router.GET("t/getbyID/:id", h.GetTID)
	router.GET("t/getAll", h.GetTAll)
	router.DELETE("t/delete/:id", h.DeleteT)

	// station api
	router.POST("/s/create", h.CreateST)
	router.PUT("/s/update/:id", h.UpdateST)
	router.GET("/s/getbyID/:id", h.GetSTID)
	router.GET("/s/getAll", h.GetSTAll)
	router.DELETE("s/delete/:id", h.DeleteST)

	return &http.Server{Handler: router, Addr: ":8075"}
}
