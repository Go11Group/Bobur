package api

import (
	"model/api_gateway/api/handler"
	genproto "model/api_gateway/genproto/prooto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	
	router := gin.Default()

	weather := genproto.NewWeatherServiceClient(conn)

	h := handler.NewHandler(weather)

	router.POST("/reportW", h.ReportWeather)
	router.GET("/currentWeather", h.GetCurrentW)
	router.GET("/weatherF", h.GetWeatherF)

	return router
}
