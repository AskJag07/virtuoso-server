package services

import (
	"encoding/json"
	"net/http"

	"github.com/AskJag07/virtuoso-server/models"
)

func ResponseJSON(message string, w http.ResponseWriter, statusCode int) {
	response := models.Response{message}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
