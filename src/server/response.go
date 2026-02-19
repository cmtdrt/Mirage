package server

import (
	"encoding/json"
	"net/http"
	"time"

	"mirage/src/models"
)

func WriteResponse(w http.ResponseWriter, ep *models.Endpoint) {
	// Apply delay if specified (in milliseconds)
	if ep.Delay != nil {
		time.Sleep(time.Duration(*ep.Delay) * time.Millisecond)
	}

	w.Header().Set("Content-Type", "application/json")
	// Set status code if specified (200 by default)
	if ep.Status != nil {
		w.WriteHeader(*ep.Status)
	}
	json.NewEncoder(w).Encode(ep.Response)
}
