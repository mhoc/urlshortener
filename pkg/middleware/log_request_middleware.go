package middleware

import (
	"log"
	"net/http"
)

// LogRequest is a simple net/http.Handler midleware for logging each request the server receives.
type LogRequest struct {
	wrap http.Handler
}

func NewLogRequest(wrap http.Handler) LogRequest {
	return LogRequest{
		wrap: wrap,
	}
}

func (m LogRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v", r.Method, r.URL.String())
	m.wrap.ServeHTTP(w, r)
}
