package job

import (
	"assignment3/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func ReadAndUpdateData() {
	currentTime := time.Now()

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	files := basePath + "/repository/data.json"

	byteValue, err := os.ReadFile(files)
	if err != nil {
		panic(err)
	}
	var data models.Data
	json.Unmarshal(byteValue, &data)

	fmt.Println("=========================================")
	fmt.Println("Updating data...")
	fmt.Println("Current Time		:", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Current Water Value	: ", data.Status.Water)
	fmt.Println("Current Wind Value	: ", data.Status.Wind)

	fmt.Println("=========================================")

	// make sure always generate random number between 1-100
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	data.Status.Water = rand.Intn(max-min) + min
	data.Status.Wind = rand.Intn(max-min) + min

	fmt.Println("Updated Water Value	: ", data.Status.Water)
	fmt.Println("Updated Wind Value	: ", data.Status.Wind)
	fmt.Println("Update data success!")

	weatherJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	os.WriteFile(files, weatherJson, 0644)
}
