package server

import (
	"encoding/json"
	"net/http"

	"mirage/src/doc"
	"mirage/src/models"
)

// Register all endpoints as HTTP handlers
func SetupRoutes(config models.Input) {
	// Built-in health check
	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	for _, ep := range config.Endpoints {
		ep := ep
		pattern := ep.Method + " " + ep.Path
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			WriteResponse(w, &ep)
		})
		doc.PrintDescription(&ep)
	}
}
