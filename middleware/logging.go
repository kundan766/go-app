package middleware

import (
	"log"
	"net/http"
)

// Logging is a middleware that logs incoming requests
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s from %s\n", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
