package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type contextKey int

const portContextKey contextKey = 0

// Returns the port number injected into the request context by the server. Returns 0 if not set.
func GetPort(r *http.Request) int {
	if r == nil || r.Context() == nil {
		return 0
	}
	v := r.Context().Value(portContextKey)
	if v == nil {
		return 0
	}
	p, _ := v.(int)
	return p
}

// Injects the port into the request context so handlers can use GetPort(r).
func wrapWithPort(port int, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), portContextKey, port)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// StartServer starts the HTTP server on a single port (blocks).
func StartServer(port int) {
	StartServers([]int{port})
}

// Starts one HTTP server per port. Each server runs in its own goroutine.
// The same routes (registered via http.HandleFunc in SetupRoutes) are served on every port.
func StartServers(ports []int) {
	if len(ports) == 0 {
		return
	}
	var wg sync.WaitGroup
	for _, port := range ports {
		port := port
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr := ":" + strconv.Itoa(port)
			fmt.Printf("Mirage running on http://localhost%s\n", addr)
			handler := wrapWithPort(port, http.DefaultServeMux)
			if err := http.ListenAndServe(addr, handler); err != nil {
				log.Printf("port %d: %v", port, err)
			}
		}()
	}
	wg.Wait()
}
