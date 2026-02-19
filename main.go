package main

import (
	"fmt"

	"mirage/src/cli"
	"mirage/src/config"
	"mirage/src/example"
	"mirage/src/server"
)

func main() {
	useExample, port, filename, err := cli.ParseFlags()
	if err != nil {
		fmt.Println("Usage: mirage serve <config.json>")
		fmt.Println("       mirage serve --example")
		fmt.Println("       mirage serve <config.json> --port=8081")
		fmt.Println("       mirage serve --example --port=8081")
		fmt.Printf("\nError: %v\n", err)
		return
	}

	// Create example file if requested
	if useExample {
		example.CreateExampleFile()
	}

	// Load configuration
	cfg := config.LoadConfig(filename)

	// Setup routes and start server
	server.SetupRoutes(cfg)
	server.StartServer(port)
}
