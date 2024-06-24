package api

import (
	"model/api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() *http.Server {
	router := gin.Default()

	h := handler.NewHandler()

	// api 8090 port un
	router.POST("/user/create", h.Create)
	router.PUT("/user/update", h.Update)
	router.DELETE("/user/delete", h.Delete)
	router.GET("/user/getbyID/:id", h.GetById)
	router.GET("/user/getAll", h.GetAll)

	// api 8075 port un
	router.POST("/create", h.Create1)
	router.PUT("/update", h.Update1)
	router.DELETE("/delete", h.Delete1)
	router.GET("/getbyID/:id", h.GetById1)
	router.GET("/getbyAll", h.GetAll1)
	


	return &http.Server{Handler: router, Addr: ":8070"}
}
