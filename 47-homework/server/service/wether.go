package service

import (
	"context"
	"log"
	"project/postgres"
	pb "project/weather/generator"
)

type weather struct {
	pb.UnimplementedWeatherServiceServer
	Weather *postgres.WeatherRepo
}

func NewWeatherSerice(db *postgres.WeatherRepo) *weather {
	return &weather{Weather: db}
}

func (w *weather) ReportWeatherCondition(cntx context.Context, req *pb.ReportWeatherRequest) (*pb.ReportWeatherReponse, error) {
	resp, err := w.Weather.ReportWeather(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *weather) GetCurrentWeather(cntx context.Context, req *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	resp, err := w.Weather.GetWeather(req)
	if err != nil {
		log.Fatal("error --> ", err.Error())
		return nil, err
	}
	return resp, nil
}

func (w *weather) GetWeatherForecast(contx context.Context, req *pb.WeatherRequest) (*pb.WeatherForecastResponse, error) {
	resps, err := w.Weather.GetWeatherForecast(req)
	if err != nil {
		return nil, err
	}
	a := 0
	for i := 0; i < len(resps); i++ {
		if resps[i].City == req.City {
			a = i
		}
	}
	return &resps[a], nil
}
