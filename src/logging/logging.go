package logging

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Represents an HTTP request handled that will be logged in /logs endpoint
type Entry struct {
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

type portState struct {
	file    *os.File
	entries []Entry
	mu      sync.Mutex
}

var (
	initMu     sync.Mutex
	portStates map[int]*portState
)

// Init initializes one log file per port.
// Files are named mirage-logs-{port}-{timestamp}.txt (e.g. mirage-logs-8080-20060102-150405.txt).
func Init(ports []int) error {
	initMu.Lock()
	defer initMu.Unlock()

	if len(ports) == 0 {
		return nil
	}

	start := time.Now()
	ts := start.Format("20060102-150405")
	portStates = make(map[int]*portState, len(ports))

	for _, port := range ports {
		filename := fmt.Sprintf("mirage-logs-%s-%s.txt", strconv.Itoa(port), ts)
		f, err := os.Create(filename)
		if err != nil {
			// close any already opened files
			for _, state := range portStates {
				_ = state.file.Close()
			}
			return err
		}
		_, _ = fmt.Fprintf(f, "START - %s (port %d)\n", start.Format(time.RFC3339), port)
		portStates[port] = &portState{file: f}
	}

	return nil
}

// LogShutdown writes a shutdown line in each port's log file and closes them.
func LogShutdown(reason string) {
	initMu.Lock()
	defer initMu.Unlock()

	for _, state := range portStates {
		state.mu.Lock()
		if state.file != nil {
			_, _ = fmt.Fprintf(state.file, "STOP  - %s - %s\n", time.Now().Format(time.RFC3339), reason)
			_ = state.file.Close()
			state.file = nil
		}
		state.mu.Unlock()
	}
	portStates = nil
}

// LogRequest appends a new entry for the given port and writes a line in that port's log file.
func LogRequest(method, path string, port int) {
	initMu.Lock()
	state, ok := portStates[port]
	initMu.Unlock()

	if !ok || state == nil {
		return
	}

	state.mu.Lock()
	defer state.mu.Unlock()

	e := Entry{
		Method:    method,
		Path:      path,
		Timestamp: time.Now(),
	}
	state.entries = append(state.entries, e)

	if state.file != nil {
		_, _ = fmt.Fprintf(
			state.file,
			"%s - %s - %s\n",
			e.Method,
			e.Timestamp.Format(time.RFC3339),
			e.Path,
		)
	}
}

// Entries returns a copy of all logged entries for the given port.
func Entries(port int) []Entry {
	initMu.Lock()
	state, ok := portStates[port]
	initMu.Unlock()

	if !ok || state == nil {
		return nil
	}

	state.mu.Lock()
	defer state.mu.Unlock()

	out := make([]Entry, len(state.entries))
	copy(out, state.entries)
	return out
}
