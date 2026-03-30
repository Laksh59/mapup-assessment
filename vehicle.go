package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Vehicle struct {
	ID            string `json:"id"`
	VehicleNumber string `json:"vehicle_number"`
	DriverName    string `json:"driver_name"`
	VehicleType   string `json:"vehicle_type"`
	Phone         string `json:"phone"`
	Status        string `json:"status"`
}

var vehicles []Vehicle

func VehicleHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// POST
	if r.Method == http.MethodPost {
		var v Vehicle

		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		v.ID = "veh_" + time.Now().Format("150405")
		v.Status = "active"

		vehicles = append(vehicles, v)

		response := map[string]interface{}{
			"id":             v.ID,
			"vehicle_number": v.VehicleNumber,
			"status":         v.Status,
			"time_ns":        time.Since(start).Nanoseconds(),
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// GET
	if r.Method == http.MethodGet {
		response := map[string]interface{}{
			"vehicles": vehicles,
			"time_ns":  time.Since(start).Nanoseconds(),
		}

		json.NewEncoder(w).Encode(response)
		return
	}
}