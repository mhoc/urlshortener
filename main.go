package main

import (
	"log"
	"net/http"
	"os"

	"github.com/twitchtv/twirp"
	"gitlab.com/mhoc/urlshortener/pkg/handler"
	"gitlab.com/mhoc/urlshortener/pkg/middleware"
	"gitlab.com/mhoc/urlshortener/pkg/proto"
)

func main() {
	log.Printf("starting server")

	// Load environment variables for configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8084"
	}

	// Create the twirp api server
	protoServer := proto.NewURLShortenerV1Server(
		proto.NewServer(),
		twirp.WithServerPathPrefix("/api"),
	)

	// Set up routes
	serveMux := http.NewServeMux()
	serveMux.Handle(protoServer.PathPrefix(), protoServer)
	// The shortlinks generated will simply look like `/{6+ random characters}` to be as short as
	// possible.
	// To handle those, this is a fall-through route which will route anything that isn't handled by
	// the longer routes defined above. Go's servemux prefers matching on longer routes.
	// One could also have a specific sub-hierarchy for links, like `/u/{id}`, but that would add
	// three additional characters to every link, and isn't really necessary given the single domain
	// focus of this service.
	serveMux.Handle("/", handler.NewShortlinkRedirect())

	// Middleware wrapping.
	wrappedMux := middleware.NewLogRequest(serveMux)

	// Start the http server
	log.Printf("serving on port %v", port)
	server := &http.Server{
		Addr:    port,
		Handler: wrappedMux,
	}
	log.Fatalf("%v", server.ListenAndServe())
}
