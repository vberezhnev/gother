package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//colorReset := "\033[0m"

	// Setting colors
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

	// Print info
	data, err := query(city)
	if err != nil {
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println()
		fmt.Println(string(colorBlue), data.Name+", "+data.Sys.Country)

		fmt.Println(string(colorYellow), "/- ====================== -/")
		fmt.Println()

		fmt.Println(string(colorGreen), "Weather (short info)", string(colorRed), data.Weather[0].Main)
		fmt.Println(string(colorGreen), "Description", string(colorRed), data.Weather[0].Description)

		fmt.Println()

		fmt.Println(string(colorGreen), "Temperature:", string(colorRed), data.Main.Celsius, "°C")
		fmt.Println(string(colorGreen), "Fells like:", string(colorRed), data.Main.FellsLike, "°C")

		fmt.Println(string(colorGreen), "Visibility:", string(colorRed), data.Visibility)
		fmt.Println(string(colorGreen), "Clouds:", string(colorRed), data.Clouds.All, "%")

		fmt.Println()
		fmt.Println(string(colorYellow), "/- ====================== -/")
		fmt.Println()
	}
}

// Extract data from API
type weatherData struct {
	Name       string  `json:"name"`
	Visibility float64 `json:"visibility"`

	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	}
	Clouds struct {
		All float64 `json:"all"`
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
	var c apiConfigData
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}

// Get data from API
func query(city string) (weatherData, error) {
	var d weatherData

	apiConfig, err := loadApiConfig(".apiConfig.json")
	if err != nil {
		return weatherData{}, errors.New("Failed to get API key.")
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city + "&units=metric")
	if err != nil {
		return weatherData{}, errors.New("Failed response from server.")
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, errors.New("Error.")
	}

	return d, nil
}
