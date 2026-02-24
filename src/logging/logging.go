package logging

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Represents an HTTP request handled that will be logged in /logs endpoint
type Entry struct {
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

var (
	mu      sync.Mutex
	entries []Entry
	logFile *os.File
)

// Creates a new log file.
// The file is named mirage-logs-<timestamp>.txt, where <timestamp> is the startup time.
func Init() error {
	mu.Lock()
	defer mu.Unlock()

	start := time.Now()
	filename := fmt.Sprintf("mirage-logs-%s.txt", start.Format("20060102-150405"))

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	logFile = f
	entries = nil

	// Write a header line with the start time.
	_, _ = fmt.Fprintf(logFile, "START - %s\n", start.Format(time.RFC3339))

	return nil
}

// Appends a new entry (for /logs endpoint) and writes a line in the file.
// One line per request: METHOD - TIMESTAMP - PATH
func LogRequest(method, path string) {
	mu.Lock()
	defer mu.Unlock()

	e := Entry{
		Method:    method,
		Path:      path,
		Timestamp: time.Now(),
	}
	entries = append(entries, e)

	if logFile != nil {
		_, _ = fmt.Fprintf(
			logFile,
			"%s - %s - %s\n",
			e.Method,
			e.Timestamp.Format(time.RFC3339),
			e.Path,
		)
	}
}

// Returns a copy of all logged entries for this run.
func Entries() []Entry {
	mu.Lock()
	defer mu.Unlock()

	out := make([]Entry, len(entries))
	copy(out, entries)
	return out
}
