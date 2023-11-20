package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)
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

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello from go!\n"))
}

func query(city string) (interface{}, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return WeatherData{}, err
	}
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return WeatherData{}, err
	}
	defer resp.Body.Close()
	var d WeatherData
	if json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return WeatherData{}, err
	}
	return d, nil
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/weather/", func(writer http.ResponseWriter, request *http.Request) {
		//以"/"分隔元素，取第三个元素，若没有则返回全部URL
		city := strings.SplitN(request.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(writer).Encode(data)
	})

	http.ListenAndServe(":9090", nil)
}
