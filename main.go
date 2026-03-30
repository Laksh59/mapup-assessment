package main

import (
	"log"
	"net/http"

	"mapup-backend/routes"
)

func main() {
	router := routes.SetupRoutes()

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}