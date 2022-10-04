package routers

import (
	"assignment3/controllers"
	"fmt"
	"net/http"
)

func StartServer() {

	// Web Page
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/monitor/", http.StripPrefix("/monitor/", http.FileServer(http.Dir("static"))))

	// API
	http.HandleFunc("/api/get-data", controllers.GetData)

	// Start Server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
