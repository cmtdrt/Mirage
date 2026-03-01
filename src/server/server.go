package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// StartServer starts the HTTP server on a single port (blocks).
func StartServer(port int) {
	StartServers([]int{port})
}

// StartServers starts one HTTP server per port. Each server runs in its own goroutine.
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
			if err := http.ListenAndServe(addr, nil); err != nil {
				log.Printf("port %d: %v", port, err)
			}
		}()
	}
	wg.Wait()
}
