package main

import (
	"log"
	"net/http"
	"os"

	"go-app/handlers"
	"go-app/middleware"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RedirectHandler)
	loggedMux := middleware.Logging(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on :" + port)
	if err := http.ListenAndServe(":"+port, loggedMux); err != nil {
		log.Fatal(err)
	}
}
