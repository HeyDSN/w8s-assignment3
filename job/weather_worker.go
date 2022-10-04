package job

import (
	"assignment3/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func ReadAndUpdateWeather() {
	currentTime := time.Now()

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

	fmt.Println("=========================================")
	fmt.Println("Updating weather...")
	fmt.Println("Current Time		:", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Current Water Value	: ", weather.Status.Water)
	fmt.Println("Current Wind Value	: ", weather.Status.Wind)

	fmt.Println("=========================================")

	// make sure always generate random number between 1-100
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	weather.Status.Water = rand.Intn(max-min) + min
	weather.Status.Wind = rand.Intn(max-min) + min

	fmt.Println("Updated Water Value	: ", weather.Status.Water)
	fmt.Println("Updated Wind Value	: ", weather.Status.Wind)
	fmt.Println("Update weather success!")

	weatherJson, err := json.Marshal(weather)
	if err != nil {
		panic(err)
	}
	os.WriteFile(files, weatherJson, 0644)
}
