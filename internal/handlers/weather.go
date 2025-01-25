package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"weather-app/internal/models"
)

type WeatherHandler struct {
	apiKey string
}

func NewWeatherHandler(apiKey string) *WeatherHandler {
	return &WeatherHandler{
		apiKey: apiKey,
	}
}

func (h *WeatherHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func (h *WeatherHandler) HandleWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "City parameter is required"})
		return
	}

	weather, err := models.FetchWeather(city, h.apiKey)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
