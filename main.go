package main

import (
	"mirage/src/cli"
	"mirage/src/config"
	"mirage/src/doc"
	"mirage/src/example"
	"mirage/src/server"
)

func main() {
	useExample, port, filename, err := cli.ParseFlags()
	if err != nil {
		doc.DisplayUsages(err)
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
