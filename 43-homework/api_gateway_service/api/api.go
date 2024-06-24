package api

import (
	"model/api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() *http.Server {
	router := gin.Default()

	h := handler.NewHandler()

	// api user
	router.POST("/user/create", h.Create)
	router.PUT("/user/update", h.Update)
	router.DELETE("/user/delete", h.Delete)
	router.GET("/user/getbyID/:id", h.GetById)
	router.GET("/user/getAll", h.GetAll)

	// api card
	router.POST("/card/create", h.Create)
	router.PUT("/card/update", h.Update)
	router.DELETE("/card/delete", h.Delete)
	router.GET("/card/getbyID/:id", h.GetById)
	router.GET("/card/getbyAll", h.GetAll)
	


	return &http.Server{Handler: router, Addr: ":8070"}
}
