package handler

import "model/api_gateway/genproto/prooto"

type Handler struct {
	Weather prooto.WeatherServiceClient
}

func NewHandler(weather prooto.WeatherServiceClient) *Handler {
	return &Handler{Weather: weather}
}
