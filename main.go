package main

import (
	"log"
	"net/http"

	"github.com/mhoc/urlshortener/pkg/config"
	"github.com/mhoc/urlshortener/pkg/handler"
	"github.com/mhoc/urlshortener/pkg/middleware"
	"github.com/mhoc/urlshortener/pkg/proto"
	"github.com/mhoc/urlshortener/pkg/store"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/twitchtv/twirp"
)

func main() {
	log.Printf("starting server")

	// Load environment variables for configuration
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Choose an appropriate datastore, and set it up
	var st store.Store
	if cfg.RedisURL != "" {
		st = store.NewRedis(cfg)
	} else {
		st = store.NewInMemoryStore()
	}

	// Register prometheus metrics
	// Of course, more metrics could be added; prometheus has a nice middleware wrapper around
	// http handlers.
	prometheus.MustRegister(handler.ShortlinkRedirectCounter)

	// Create the twirp api server
	protoServer := proto.NewURLShortenerV1Server(
		proto.NewServer(cfg, st),
		twirp.WithServerPathPrefix("/api"),
	)

	// Set up routes
	serveMux := http.NewServeMux()
	serveMux.Handle(protoServer.PathPrefix(), protoServer)
	serveMux.Handle("/metrics", promhttp.Handler())
	// The shortlinks generated will simply look like `/{8 random characters}` to be as short as
	// possible.
	// To handle those, this is a fall-through route which will route anything that isn't handled by
	// the longer routes defined above. Go's servemux prefers matching on longer routes.
	// One could also have a specific sub-hierarchy for links, like `/u/{id}`, but that would add
	// three additional characters to every link, and isn't really necessary given the single domain
	// focus of this service.
	serveMux.Handle("/", handler.NewShortlinkRedirect(st))

	// Middleware wrapping.
	wrappedMux := middleware.NewLogRequest(serveMux)

	// Start the http server
	log.Printf("serving on port %v", cfg.Port)
	server := &http.Server{
		Addr:    cfg.Port,
		Handler: wrappedMux,
	}
	log.Fatalf("%v", server.ListenAndServe())
}
