package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64 `json:"time_epoch"`
				TempC float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}



func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == ""{
		fmt.Println("Error, WEATHER_API_KEY not found")
		os.Exit(1)
	}
	loca := "Mumbai"
	if len(os.Args) > 1 {
		loca = os.Args[1]
	}
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, loca))
	if err != nil {
		panic("Error: %s" +err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Error, Unable to fetch weather data")
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Error: %s" +err.Error())
	}
	// fmt.Println(string(body))

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err.Error()))
	}
	// fmt.Println(weather)

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf("Location: %s, %s\n", location.Name, location.Country)
	fmt.Printf("Current Temperature: %.2f C\n", current.TempC)
	fmt.Printf("Condition: %s\n", current.Condition.Text)
	fmt.Println("Hourly Forecast:")
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		now := time.Now()
		midnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		if date.Before(now) || date.After(midnight) {
			continue
		}else{
		color.Red("Time: %s, %.2f C, %s, Chance of Rain: %.2f\n", date, hour.TempC, hour.Condition.Text, hour.ChanceOfRain)
		}
	}
	

}