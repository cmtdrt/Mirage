package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"mirage/src/models"
)

func LoadConfig(filename string) models.Input {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filename, err)
	}

	var config models.Input
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	fmt.Println("\n================ MIRAGE ================")
	fmt.Printf("\nFound %d endpoint(s)\n", len(config.Endpoints))
	fmt.Println("")

	return config
}
