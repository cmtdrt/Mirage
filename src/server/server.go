package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// StartServer starts the HTTP server on the specified port
func StartServer(port int) {
	addr := ":" + strconv.Itoa(port)
	fmt.Printf("\nMirage running on port %d\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
