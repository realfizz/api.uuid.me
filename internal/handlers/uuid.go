package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"api.uuid.me/internal/models"
)


func UUIDHandler(w http.ResponseWritter, r *http.Request) {
	newUUID := uuid.new() // default is v4 btw

	json.NewEncoder(w).Encode(models.Message{
		Message: newUUID.String(),
	}
}
