package controllers

import (
	"assignment3/models"
	"encoding/json"
	"net/http"
	"os"
)

func GetData(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
