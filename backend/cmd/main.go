package main

import (
	"log"
	"net/http"
)

// A simple web server with a health check endpoint
func main() {
	// Health check endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Health check endpoint hit")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
