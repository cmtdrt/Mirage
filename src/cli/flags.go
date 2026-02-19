package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	EXAMPLE_FLAG      = "--example"
	PORT_FLAG         = "--port="
	DEFAULT_PORT      = 8080
	EXAMPLE_FILE_NAME = "mirage.example.json"
	DEFAULT_FILE_NAME = "mirage.json"
)

// ParseFlags parses command line flags and returns the configuration
func ParseFlags() (useExample bool, port int, filename string, err error) {
	port = DEFAULT_PORT
	useExample = false

	if len(os.Args) < 3 {
		return false, 0, "", fmt.Errorf("insufficient arguments")
	}

	// Parse all arguments starting from index 2
	for i := 2; i < len(os.Args); i++ {
		arg := os.Args[i]

		if arg == EXAMPLE_FLAG {
			useExample = true
		} else if strings.HasPrefix(arg, PORT_FLAG) {
			portStr := strings.TrimPrefix(arg, PORT_FLAG)
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
		filename = EXAMPLE_FILE_NAME
	}

	// If no filename and not using example, try to find mirage.json
	if filename == "" && !useExample {
		if _, err := os.Stat(DEFAULT_FILE_NAME); err == nil {
			filename = DEFAULT_FILE_NAME
		} else {
			return false, 0, "", fmt.Errorf("no config file specified and mirage.json not found")
		}
	}

	return useExample, port, filename, nil
}
