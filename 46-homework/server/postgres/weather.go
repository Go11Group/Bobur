package postgres

import (
	"database/sql"
	pb "project/weather/generator"
)

type WeatherRepo struct {
	Db *sql.DB
}

func NewWeatherRepo(db *sql.DB) *WeatherRepo {
	return &WeatherRepo{db}
}

func (w *WeatherRepo) GetWeather(req *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	resp := pb.WeatherResponse{}
	row := w.Db.QueryRow("select city, temperature, description from weather2 where city=$1", req.City)
	err := row.Scan(&resp.City, &resp.Temperature, &resp.Description)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (w *WeatherRepo) ReportWeather(req *pb.ReportWeatherRequest) (*pb.ReportWeatherReponse, error) {
	_, err := w.Db.Exec("inser into weather1(name, date, windS, tempernature, rain, sun) values($1, $2, $3, $4, $5, $6)",
		req.Name, req.Date, req.WindS, req.Tempernature, req.Rain, req.Sun)
	if err != nil {
		return nil, err
	}
	return &pb.ReportWeatherReponse{Query: "SUCCESS"}, nil
}

func (w *WeatherRepo) GetWeatherForecast(req *pb.WeatherRequest) ([]pb.WeatherForecastResponse, error) {
	rows, err := w.Db.Query("select city, date, temperature, description from weather3 where city=$1", req.City)
	if err != nil {
		return nil, err
	}

	resps := []pb.WeatherForecastResponse{}
	i := 0
	for rows.Next() {
		reps := pb.WeatherForecastResponse{}
		err := rows.Scan(&reps.City, &reps.DailyForecasts[i].Date, &reps.DailyForecasts[i].Temperature, &reps.DailyForecasts[i].Description)
		if err != nil {
			return nil, err
		}

		resps = append(resps, resps...)
	}

	return resps, nil
}
