package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type LocationRequest struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp string  `json:"timestamp"`
}

// Store previous state
var vehicleGeofenceState = make(map[string]string)

// 🔥 Point in polygon
func isPointInPolygon(lat, lon float64, polygon [][]float64) bool {
	inside := false
	j := len(polygon) - 1

	for i := 0; i < len(polygon); i++ {
		xi, yi := polygon[i][0], polygon[i][1]
		xj, yj := polygon[j][0], polygon[j][1]

		intersect := ((yi > lon) != (yj > lon)) &&
			(lat < (xj-xi)*(lon-yi)/(yj-yi)+xi)

		if intersect {
			inside = !inside
		}
		j = i
	}

	return inside
}

func VehicleLocationHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.Method == http.MethodPost {
		var loc LocationRequest

		err := json.NewDecoder(r.Body).Decode(&loc)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		var currentGeofenceID string
		var currentGeofences []map[string]interface{}

		for _, geo := range geofences {
			if isPointInPolygon(loc.Latitude, loc.Longitude, geo.Coordinates) {
				currentGeofenceID = geo.ID

				currentGeofences = append(currentGeofences, map[string]interface{}{
					"geofence_id":   geo.ID,
					"geofence_name": geo.Name,
					"status":        "inside",
				})
			}
		}

		// Previous state
		prev := vehicleGeofenceState[loc.VehicleID]

		var event string

		if prev == "" && currentGeofenceID != "" {
			event = "ENTRY"
		} else if prev != "" && currentGeofenceID == "" {
			event = "EXIT"
		}

		// Update state
		vehicleGeofenceState[loc.VehicleID] = currentGeofenceID

		response := map[string]interface{}{
			"vehicle_id":        loc.VehicleID,
			"location_updated":  true,
			"current_geofences": currentGeofences,
			"event":             event,
			"time_ns":           time.Since(start).Nanoseconds(),
		}

		json.NewEncoder(w).Encode(response)
	}
}