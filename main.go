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
	useExample, ports, filename, err := cli.ParseFlags()
	if err != nil {
		doc.DisplayUsages(err)
		return
	}

	// Initialize logging (one file per port: mirage-logs-{port}-{timestamp}.txt)
	if err := logging.Init(ports); err != nil {
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

	// Setup routes and start server(s)
	server.SetupRoutes(cfg)
	server.StartServers(ports)
}
