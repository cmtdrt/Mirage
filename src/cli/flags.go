package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ExampleFlag = "--example"
	PortFlag    = "--port="
)

// ParseFlags parses command line flags and returns the configuration
func ParseFlags() (useExample bool, port int, filename string, err error) {
	port = 8080 // default port
	useExample = false

	if len(os.Args) < 3 {
		return false, 0, "", fmt.Errorf("insufficient arguments")
	}

	// Parse all arguments starting from index 2
	for i := 2; i < len(os.Args); i++ {
		arg := os.Args[i]

		if arg == ExampleFlag {
			useExample = true
		} else if strings.HasPrefix(arg, PortFlag) {
			portStr := strings.TrimPrefix(arg, PortFlag)
			parsedPort, err := strconv.Atoi(portStr)
			if err != nil {
				return false, 0, "", fmt.Errorf("invalid port value: %s", portStr)
			}
			if parsedPort < 1 || parsedPort > 65535 {
				return false, 0, "", fmt.Errorf("port must be between 1 and 65535")
			}
			port = parsedPort
		} else if !strings.HasPrefix(arg, "--") {
			// This is the filename (not a flag)
			if filename != "" {
				return false, 0, "", fmt.Errorf("multiple filenames specified")
			}
			filename = arg
		}
	}

	// If --example is used but no filename, use the example file
	if useExample && filename == "" {
		filename = "mirage.example.json"
	}

	// If no filename and not using example, try to find mirage.json
	if filename == "" && !useExample {
		if _, err := os.Stat("mirage.json"); err == nil {
			filename = "mirage.json"
		} else {
			return false, 0, "", fmt.Errorf("no config file specified and mirage.json not found")
		}
	}

	return useExample, port, filename, nil
}
