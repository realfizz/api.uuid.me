package handlers

import (
	"encoding/json"
	"net/http"

	"api.uuid.me/internal/models"
)


func HealthHandler(w http.ResponseWriter, r *http.Request){ 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Message{
		Message: "The mainframe is running",
	})
}


