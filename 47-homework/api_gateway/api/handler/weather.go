package handler

import (
	pb "model/api_gateway/genproto/prooto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ReportWeather(g *gin.Context) {
	req := pb.ReportWeatherRequest{}

	err := g.ShouldBindJSON(&req)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Weather.ReportWeatherCondition(g, &req)
	if err != nil {
		
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCurrentW(g *gin.Context) {
	req := pb.WeatherRequest{}
	resp, err := h.Weather.GetCurrentWeather(g, &req)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, resp)
}

func (h *Handler) GetWeatherF(g *gin.Context) {
	req := pb.WeatherRequest{}

	resp, err := h.Weather.GetWeatherForecast(g, &req)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, resp)
}
