package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	w "model/genproto/prooto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func clear() {
	// Clearing the screen based on the operating system
	switch runtime.GOOS {
	case "darwin", "linux", "freebsd", "openbsd", "netbsd":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported OS")
	}
}
func displayMenu(client w.WeatherServiceClient) {
	for {
		// clear()
		fmt.Println("1. Report Weather Condition")
		fmt.Println("2. Get Current weather")
		fmt.Println("3. Get Weather Forecast")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			clear()
			ReportWeatherCondition(client)
		case 2:
			clear()
			GetCurrentWeatherCondition(&client)
			fmt.Scan(&choice)
		case 3:
			GetWeatherForecast(&client)
			fmt.Scan(&choice)

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func main() {
	conn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := w.NewWeatherServiceClient(conn)
	clear()
	fmt.Println("Welcome to Weather service")

	displayMenu(client)

}

func ReportWeatherCondition(client w.WeatherServiceClient) {
	wCon := w.ReportWeatherRequest{}
	fmt.Println("Enter the location:")
	fmt.Scanln(&wCon.Name)
	fmt.Println("Enter the temperature:")
	fmt.Scanln(&wCon.Tempernature)
	fmt.Println("Enter the wind speed:")
	fmt.Scanln(&wCon.WindS)
	fmt.Println("Enter the rain:")
	fmt.Scanln(&wCon.Rain)
	fmt.Println("Enter the sun:")
	fmt.Scanln(&wCon.Sun)

	res, err := client.ReportWeatherCondition(context.Background(), &wCon)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Weather condition has been reported: %w", res)
}

func GetCurrentWeatherCondition(client *w.WeatherServiceClient) {
	fmt.Println("Enter the location:")
	var City string
	fmt.Scanln(&City)

	resp, err := (*client).GetCurrentWeather(context.Background(), &w.WeatherRequest{City: City})
	if err != nil {
		log.Println("getting current weather failed: ", err)
		return
	}

	fmt.Println("Current weather in: ", City)
	fmt.Println("Temperature: ", resp.Temperature)
	fmt.Println("Desctiption: ", resp.Description)

}

func GetWeatherForecast(clinet *w.WeatherServiceClient) {
	fmt.Println("Enter the location:")
	var City string
	fmt.Scanln(&City)
	resp, err := (*clinet).GetWeatherForecast(context.Background(), &w.WeatherRequest{City: City})
	if err != nil {
		log.Println("getting current weather failed: ", err)
		return
	}
	fmt.Println("Weather forecast in: ", City)
	fmt.Println(resp.DailyForecasts)
}
