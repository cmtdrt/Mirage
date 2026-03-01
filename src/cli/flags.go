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
	PORTS_FLAG        = "--ports="
	DEFAULT_PORT      = 8080
	EXAMPLE_FILE_NAME = "mirage.example.json"
	DEFAULT_FILE_NAME = "mirage.json"
)

// ParseFlags parses command line flags for the "serve" command.
// Returns ports as a slice: one element for --port=X, multiple for --ports=8080,8081,8082.
func ParseFlags() (useExample bool, ports []int, filename string, err error) {
	ports = []int{DEFAULT_PORT}
	useExample = false

	if len(os.Args) < 3 {
		return false, nil, "", fmt.Errorf("insufficient arguments")
	}

	if os.Args[1] != "serve" {
		return false, nil, "", fmt.Errorf("unknown command: %s", os.Args[1])
	}

	// Parse all arguments starting from index 2
	for i := 2; i < len(os.Args); i++ {
		arg := os.Args[i]

		if arg == EXAMPLE_FLAG {
			useExample = true
		} else if strings.HasPrefix(arg, PORTS_FLAG) {
			portsStr := strings.TrimPrefix(arg, PORTS_FLAG)
			parsed, parseErr := parsePortList(portsStr)
			if parseErr != nil {
				return false, nil, "", parseErr
			}
			ports = parsed
		} else if strings.HasPrefix(arg, PORT_FLAG) {
			portStr := strings.TrimPrefix(arg, PORT_FLAG)
			parsedPort, parseErr := strconv.Atoi(portStr)
			if parseErr != nil {
				return false, nil, "", fmt.Errorf("invalid port value: %s", portStr)
			}
			if parsedPort < 1 || parsedPort > 65535 {
				return false, nil, "", fmt.Errorf("port must be between 1 and 65535")
			}
			ports = []int{parsedPort}
		} else if !strings.HasPrefix(arg, "--") {
			// This is the filename (not a flag)
			if filename != "" {
				return false, nil, "", fmt.Errorf("multiple filenames specified")
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
		if _, statErr := os.Stat(DEFAULT_FILE_NAME); statErr == nil {
			filename = DEFAULT_FILE_NAME
		} else {
			return false, nil, "", fmt.Errorf("no config file specified and mirage.json not found")
		}
	}

	return useExample, ports, filename, nil
}

// parsePortList parses a comma-separated list of ports (e.g. "8080,8081,8082").
func parsePortList(s string) ([]int, error) {
	if s == "" {
		return nil, fmt.Errorf("--ports= requires at least one port")
	}
	parts := strings.Split(s, ",")
	ports := make([]int, 0, len(parts))
	seen := make(map[int]bool)
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		num, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("invalid port in --ports=: %s", p)
		}
		if num < 1 || num > 65535 {
			return nil, fmt.Errorf("port must be between 1 and 65535, got %d", num)
		}
		if seen[num] {
			continue // skip duplicates
		}
		seen[num] = true
		ports = append(ports, num)
	}
	if len(ports) == 0 {
		return nil, fmt.Errorf("--ports= requires at least one port")
	}
	return ports, nil
}
