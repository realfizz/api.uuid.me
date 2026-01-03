package routes

import (
	"net/http"

	"api.uuid.me/internal/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/uuid", handlers.UUIDHandler)
}
