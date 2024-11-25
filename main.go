package main

import (
	"log"
	"net/http"

	"go-app/handlers"
	"go-app/middleware"
)

func main() {
	// Create a new HTTP server
	mux := http.NewServeMux()

	// Register the redirect handler
	mux.HandleFunc("/", handlers.RedirectHandler)

	// Wrap the mux with logging middleware
	loggedMux := middleware.Logging(mux)

	// Start the server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatal(err)
	}
}
