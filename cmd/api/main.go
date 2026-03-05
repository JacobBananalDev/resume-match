package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

// main - entrypoint for the ResumeMatch API server.
// Keep main.go small: it shall mainly wire up config + routes + start the server
// Business logic will live elsewhere later (eg. internal/packages)
func main(){
	// Read port from enviornment variables so this works locally and in Docker/cloud
	// Default to 8080 for local dev
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	// Create a new chi router 
	// chi - lightweight and idiomatic for GO apis
	r := chi.NewRouter()

	// A simple health check endpoint 
	// Useful for: "Is the server alive?", docker health checks, and load balancers
	r.Get("/health", func(w http.ResponseWriter, r *http.Request){
		// Return plain text "ok" with 200 status
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

		// Create an HTTP server with basic safety timeouts.
	// Timeouts protect you from slowloris-style attacks and hung connections.
	srv := &http.Server{
		Addr: ":" + port,

		// Chi router handles all requests.
		Handler: r,

		// Basic production-friendly timeouts (still simple).
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("ResumeMatch API listening on http://localhost:%s", port)

	// Start the server (blocking call).
	// If it returns an error, log.Fatal will print it and exit.
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}