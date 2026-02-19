package main

import (
	"fmt"
	"os"

	"mirage/src/config"
	"mirage/src/example"
	"mirage/src/server"
)

func main() {
	var filename string

	// Check for --example flag
	if len(os.Args) >= 3 && os.Args[2] == example.ExampleFlag {
		example.CreateExampleFile()
		filename = example.OutputFileName
	} else {
		if len(os.Args) < 3 {
			fmt.Println("Usage: mirage serve <config.json>")
			fmt.Println("       mirage serve --example")
			return
		}
		filename = os.Args[2]
	}

	// Load configuration
	cfg := config.LoadConfig(filename)

	// Setup routes and start server
	server.SetupRoutes(cfg)
	server.StartServer()
}
