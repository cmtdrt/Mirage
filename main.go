package main

import (
	"mirage/src/cli"
	"mirage/src/config"
	"mirage/src/doc"
	"mirage/src/example"
	"mirage/src/logging"
	"mirage/src/server"
	"os"
)

func main() {
	// Commands that only generate a file and exit
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "guide-en":
			doc.GenerateGuide("en")
			return
		case "guide-fr":
			doc.GenerateGuide("fr")
			return
		}
	}

	// serve command
	useExample, port, filename, err := cli.ParseFlags()
	if err != nil {
		doc.DisplayUsages(err)
		return
	}

	// Initialize logging
	if err := logging.Init(); err != nil {
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
