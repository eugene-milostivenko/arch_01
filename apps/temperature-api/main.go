package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

// [-10.0, 40.0]
func randomTemperature() float64 {
	t := -10.0 + rand.Float64()*50.0
	return float64(int(t*10)) / 10
}

func newResponse(location, sensorID string) TemperatureResponse {
	return TemperatureResponse{
		Value:       randomTemperature(),
		Unit:        "°C",
		Timestamp:   time.Now().UTC(),
		Location:    location,
		Status:      "active",
		SensorID:    sensorID,
		SensorType:  "temperature",
		Description: "Simulated temperature reading",
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func handleTemperature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	// ID: /temperature/<sensorID>
	if id := strings.TrimPrefix(r.URL.Path, "/temperature/"); id != "" && id != r.URL.Path {
		resp := newResponse("sensor-"+id, id)
		log.Printf("GET /temperature/%s -> %.1f%s", id, resp.Value, resp.Unit)
		writeJSON(w, http.StatusOK, resp)
		return
	}

	// location: /temperature?location=<name>
	location := r.URL.Query().Get("location")
	if location == "" {
		location = "unknown"
	}
	resp := newResponse(location, "")
	log.Printf("GET /temperature?location=%s -> %.1f%s", location, resp.Value, resp.Unit)
	writeJSON(w, http.StatusOK, resp)
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/temperature", handleTemperature)
	mux.HandleFunc("/temperature/", handleTemperature)
	mux.HandleFunc("/health", handleHealth)

	addr := ":" + port
	log.Printf("temperature-api listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
