package controllers

import (
	"assignment3/models"
	"encoding/json"
	"net/http"
	"os"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files := basePath + "/job/weather.json"

	byteValue, err := os.ReadFile(files)
	if err != nil {
		panic(err)
	}
	var weather models.Weather
	json.Unmarshal(byteValue, &weather)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weather)
}
