package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	// colorCyan := "\033[36m"
	//colorWhite := "\033[37m"

	var city string

	fmt.Print(string(colorPurple), "Type a city >>> ")
	fmt.Scanln(&city)

	data, err := query(city)
	if err != nil {
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}

	if city != "" {
		fmt.Println(string(colorBlue), data.Name, ",", data.Sys.Country)

		fmt.Println()
		fmt.Println(string(colorYellow), "/- ====================== -/")
		fmt.Println()

		fmt.Println(string(colorGreen), "Weather (short info)", string(colorRed), data.Weather)
		fmt.Println(string(colorGreen), "Description", string(colorRed), data.Weather)

		fmt.Println()

		fmt.Println(string(colorGreen), "Temperature", "is", string(colorRed), data.Main.Celsius, "°C", string(colorReset))
		fmt.Println(string(colorGreen), "Fells like", string(colorRed), data.Main.FellsLike)

		fmt.Println()
		fmt.Println(string(colorYellow), "/- ====================== -/")
		fmt.Println()
	} else {
		fmt.Println(string(colorRed), "Nope")
	}
}

// Extract data from API
type weatherData struct {
	Name    string `json:"name"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	}
	Sys struct {
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

// Use API from file
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
  var d weatherData
	
  /* apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	} */

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=0fc50e5fe3acdddbea0f8fd6d9795f8b&q=" + city + "&units=metric")
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}
