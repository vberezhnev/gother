package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"
)

func main() {
		var city string
	
    fmt.Print("Type a city >>> ")
		fmt.Scanln(&city)
		fmt.Println("You type: ", city)

		data, err := query(city)
		if err != nil {
			fmt.Println(err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(data)
}

// Extract data from API
type weatherData struct {
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	}
	Main struct {
		Celsius   float64 `json:"temp"`
		FellsLike float64 `json:"feels_like"`
	} `json:"main"`
}

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}

func query(city string) (weatherData, error) {
	/* apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	} */

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=0fc50e5fe3acdddbea0f8fd6d9795f8b&q=" + city + "&units=metric")
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}
