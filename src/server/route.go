package server

import (
	"net/http"

	"mirage/src/doc"
	"mirage/src/models"
)

// Register all endpoints as HTTP handlers
func SetupRoutes(config models.Input) {
	for _, ep := range config.Endpoints {
		ep := ep
		pattern := ep.Method + " " + ep.Path
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			WriteResponse(w, &ep)
		})
		doc.PrintDescription(&ep)
	}
}
