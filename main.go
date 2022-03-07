package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/mhoc/urlshortener/pkg/handlers"
)

func main() {
	log.Printf("starting server")
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8084"
	}

	serveMux := http.NewServeMux()
	serveMux.Handle("/healthz", handlers.NewHealthCheckHandler())
	serveMux.Handle("/shortlink/create", handlers.NewShortlinkCreateHandler())
	serveMux.Handle("/shortlink/remove", handlers.NewShortlinkRemoveHandler())
	serveMux.Handle("/", handlers.NewShortlinkRedirectHandler())

	log.Printf("serving on port %v", port)
	server := &http.Server{
		Addr: port,
	}
	log.Fatalf("%v", server.ListenAndServe())
}
