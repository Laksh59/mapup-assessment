package routes

import (
	"net/http"

	"mapup-backend/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/geofences", handlers.GeofenceHandler)
	mux.HandleFunc("/vehicles", handlers.VehicleHandler)
	mux.HandleFunc("/vehicles/location", handlers.VehicleLocationHandler)

	return mux
}
