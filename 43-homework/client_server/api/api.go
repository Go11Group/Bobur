package api

import (
	"database/sql"
	"model/api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	router := gin.Default()

	h := handler.NewHandler(db)

	router.POST("user/create", h.Create)
	router.PUT("user/update/:id", h.Update)
	router.GET("user/getbyID/:id", h.GetById)
	router.GET("user/getAll", h.GetAll)
	router.DELETE("user/delete/:id", h.Delete)

	return &http.Server{Handler: router, Addr: ":8090"}
}