package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Content of the JSON file
type Input struct {
	Endpoints []Endpoint `json:"endpoints"`
}

// Object that will become a route
type Endpoint struct {
	Method      string  `json:"method"`
	Description *string `json:"description,omitempty"`
	Path        string  `json:"path"`
	Status      *int    `json:"status,omitempty"`
	Response    any     `json:"response"`
}

func main() {
	// Process the input and return the config
	config := processInput()

	// For each endpoint in config, create a route
	for _, ep := range config.Endpoints {
		pattern := ep.Method + " " + ep.Path
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			writeResponse(w, ep)
		})
		writeDescription(ep)
	}

	fmt.Println("\nMirage running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func processInput() Input {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mirage serve <config.json>")
		return Input{}
	}

	// Read the JSON file
	filename := os.Args[2]
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filename, err)
	}

	var config Input
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	fmt.Println("\n================ MIRAGE ================")
	fmt.Println("\nFound", len(config.Endpoints), "endpoint(s)")
	fmt.Println("")

	return config
}

func writeResponse(w http.ResponseWriter, ep Endpoint) {
	w.Header().Set("Content-Type", "application/json")
	if ep.Status != nil {
		w.WriteHeader(*ep.Status)
	}
	json.NewEncoder(w).Encode(ep.Response)
}

func writeDescription(ep Endpoint) {
	desc := ""
	if ep.Description != nil {
		desc = "-> " + *ep.Description
	}
	fmt.Printf("%s '%s' %s\n", ep.Method, ep.Path, desc)
}
