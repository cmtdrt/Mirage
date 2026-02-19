package example

import (
	_ "embed"
	"fmt"
	"log"
	"os"
)

//go:embed mirage-example-content.json
var exampleContent []byte

const OutputFileName = "mirage.example.json"
const ExampleFlag = "--example"

// CreateExampleFile creates the example configuration file from the embedded template
func CreateExampleFile() {
	err := os.WriteFile(OutputFileName, exampleContent, 0644)
	if err != nil {
		log.Fatalf("Failed to create example file: %v", err)
	}
	fmt.Printf("Created example file: %s\n", OutputFileName)
}
