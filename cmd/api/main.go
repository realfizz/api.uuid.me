package main

import (
	"log"
	"net/http"

	"api.uuid.me/internal/routes"
)

func main() {
	routes.RegisterRoutes()

	log.Println("Mainframe is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

