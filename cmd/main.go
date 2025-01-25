package main

import (
	"log"
	"net/http"
	"os"
	"weather-app/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new weather handler
	weatherHandler := handlers.NewWeatherHandler(os.Getenv("OPENWEATHER_API_KEY"))

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle routes
	http.HandleFunc("/", weatherHandler.HandleHome)
	http.HandleFunc("/weather", weatherHandler.HandleWeather)

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
} 