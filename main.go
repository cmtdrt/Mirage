package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const EXAMPLE_FILE_NAME = "mirage.example.json"

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
	Delay       *int    `json:"delay,omitempty"`
	Response    any     `json:"response"`
}

func main() {
	var filename string

	// Check for --example flag
	if len(os.Args) >= 3 && os.Args[2] == "--example" {
		createExampleFile()
		filename = EXAMPLE_FILE_NAME
	} else {
		if len(os.Args) < 3 {
			fmt.Println("Usage: mirage serve <config.json>")
			fmt.Println("       mirage serve --example")
			return
		}
		filename = os.Args[2]
	}

	// Process the input and return the config
	config := processInput(filename)

	for _, ep := range config.Endpoints {
		ep := ep
		pattern := ep.Method + " " + ep.Path
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			writeResponse(w, ep)
		})
		writeDescription(ep)
	}

	fmt.Println("\nMirage running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createExampleFile() {
	exampleContent := `
	{
    "endpoints": [
        {
            "method": "GET",
            "description": "Just saying hello",
            "path": "/hello",
            "delay": 1000,
            "response": "Hi there 👋"
        },
        {
            "method": "GET",
            "path": "/api/v1/users",
            "response": {
            "users": [
                {
                    "id": 1,
                    "username": "cmtdrt",
                    "email": "cmtdrt@example.com",
                    "firstName": "Clément",
                    "lastName": "Drt",
                    "role": "ADMIN",
                    "isActive": true
                }
            ]
            }
        },
        {
            "method": "GET",
            "path": "/api/v1/users/{id}",
            "response": {
                "id": 1,
                "username": "cmtdrt",
                "email": "cmtdrt@example.com"
            }
        }
    ]
  }`

	filename := "mirage.example.json"
	err := os.WriteFile(filename, []byte(exampleContent), 0644)
	if err != nil {
		log.Fatalf("Failed to create example file: %v", err)
	}
	fmt.Printf("Created example file: %s\n", filename)
}

func processInput(filename string) Input {
	// Read the JSON file
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
	// Apply delay if specified (in milliseconds)
	if ep.Delay != nil {
		time.Sleep(time.Duration(*ep.Delay) * time.Millisecond)
	}

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
