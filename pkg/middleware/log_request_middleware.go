package middleware

import (
	"log"
	"net/http"
)

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
