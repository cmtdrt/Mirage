package server

import (
	"encoding/json"
	"net/http"

	"mirage/src/doc"
	"mirage/src/logging"
	"mirage/src/models"
)

const (
	GET_REQUEST_METHOD = "GET"
	CONTENT_TYPE_KEY   = "Content-Type"
	CONTENT_TYPE_JSON  = "application/json"
)

// Register all endpoints as HTTP handlers
func SetupRoutes(config models.Input) {
	// Built-in health check
	http.HandleFunc(GET_REQUEST_METHOD+" /health", func(w http.ResponseWriter, r *http.Request) {
		logging.LogRequest(r.Method, r.URL.Path, GetPort(r))
		w.Header().Set(CONTENT_TYPE_KEY, CONTENT_TYPE_JSON)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Logs endpoint: returns all logged calls as JSON for this port
	http.HandleFunc(GET_REQUEST_METHOD+" /logs", func(w http.ResponseWriter, r *http.Request) {
		logging.LogRequest(r.Method, r.URL.Path, GetPort(r))
		w.Header().Set(CONTENT_TYPE_KEY, CONTENT_TYPE_JSON)
		json.NewEncoder(w).Encode(logging.Entries(GetPort(r)))
	})

	for _, ep := range config.Endpoints {
		ep := ep
		pattern := ep.Method + " " + ep.Path
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			logging.LogRequest(r.Method, r.URL.Path, GetPort(r))
			WriteResponse(w, &ep, r)
		})
		doc.PrintDescription(&ep)
	}
}
