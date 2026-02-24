package main

import (
	"os"
	"os/signal"
	"syscall"

	"mirage/src/cli"
	"mirage/src/config"
	"mirage/src/doc"
	"mirage/src/example"
	"mirage/src/logging"
	"mirage/src/server"
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

	// Handle shutdown signals to log when the API is stopped (Ctrl+C, SIGTERM, etc.).
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		logging.LogShutdown(sig.String())
		os.Exit(0)
	}()

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
