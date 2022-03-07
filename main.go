package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/mhoc/urlshortener/pkg/handler"
	"gitlab.com/mhoc/urlshortener/pkg/middleware"
)

func main() {
	log.Printf("starting server")

	// Load environment variables for configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8084"
	}

	// Set up routes
	serveMux := http.NewServeMux()
	serveMux.Handle("/healthz", handler.NewHealthCheck())
	serveMux.Handle("/shortlink/create", handler.NewShortlinkCreate())
	serveMux.Handle("/shortlink/remove", handler.NewShortlinkRemove())
	serveMux.Handle("/", handler.NewShortlinkRedirect())

	// Wrap middleware
	wrappedMux := middleware.NewLogRequest(serveMux)

	// Start the http server
	log.Printf("serving on port %v", port)
	server := &http.Server{
		Addr:    port,
		Handler: wrappedMux,
	}
	log.Fatalf("%v", server.ListenAndServe())
}
