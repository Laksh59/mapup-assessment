package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Geofence struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Coordinates [][]float64   `json:"coordinates"`
	Category    string        `json:"category"`
}

var geofences []Geofence

func GeofenceHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// POST
	if r.Method == http.MethodPost {
		var geo Geofence

		err := json.NewDecoder(r.Body).Decode(&geo)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		geo.ID = "geo_" + time.Now().Format("150405")
		geofences = append(geofences, geo)

		response := map[string]interface{}{
			"id":      geo.ID,
			"name":    geo.Name,
			"status":  "active",
			"time_ns": time.Since(start).Nanoseconds(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// GET
	if r.Method == http.MethodGet {
		response := map[string]interface{}{
			"geofences": geofences,
			"time_ns":   time.Since(start).Nanoseconds(),
		}

		json.NewEncoder(w).Encode(response)
		return
	}
}
