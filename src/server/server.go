package server

import (
	"fmt"
	"log"
	"net/http"
)

const DefaultPort = ":8080"

func StartServer() {
	fmt.Println("\nMirage running on port 8080")
	log.Fatal(http.ListenAndServe(DefaultPort, nil))
}
