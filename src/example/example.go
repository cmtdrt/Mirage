package example

import (
	_ "embed"
	"fmt"
	"log"
	"os"
)

//go:embed mirage-example-content.json
var exampleContent []byte

const CREATED_EXAMPLE_FILE_NAME = "mirage.example.json"

func CreateExampleFile() {
	err := os.WriteFile(CREATED_EXAMPLE_FILE_NAME, exampleContent, 0644)
	if err != nil {
		log.Fatalf("Failed to create example file: %v", err)
	}
	fmt.Printf("Created example file: %s\n", CREATED_EXAMPLE_FILE_NAME)
}
