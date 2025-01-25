package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Struct to hold API key configuration
type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

// Struct to hold weather data response
type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

// Load API key from config file

func loadApiConfig(filename string) (apiConfigData, error) {
	fmt.Println("Looking for file:", filename) // Add this for debugging
	bytes, err := ioutil.ReadFile(filename)
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

// Simple hello handler
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go!\n"))
}

// Function to query weather data from OpenWeatherMap API
func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig("apiConfig.json") // Use the correct file name
	if err != nil {
		return weatherData{}, err
	}

	// Construct the API request URL correctly
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s&units=metric", apiConfig.OpenWeatherMapApiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	// Handle API response errors
	if resp.StatusCode != http.StatusOK {
		return weatherData{}, fmt.Errorf("failed to fetch weather data: %s", resp.Status)
	}

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}

func main() {
	// Hello route
	http.HandleFunc("/hello", hello)

	// Weather route
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.SplitN(r.URL.Path, "/", 3)
		if len(parts) < 3 || strings.TrimSpace(parts[2]) == "" {
			http.Error(w, "City name is required", http.StatusBadRequest)
			return
		}

		city := parts[2]
		log.Println("Fetching weather for city:", city)

		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set proper content-type
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	// Start the server
	log.Println("Starting server on port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
