package routers

import (
	"assignment3/controllers"
	"fmt"
	"net/http"
)

func StartServer() {

	// Web Page
	http.Handle("/weather/", http.StripPrefix("/weather/", http.FileServer(http.Dir("static"))))

	// API
	http.HandleFunc("/api/weather", controllers.GetWeather)

	// Start Server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
