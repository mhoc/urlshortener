package handler

import "net/http"

type HealthCheck struct{}

func NewHealthCheck() HealthCheck {
	return HealthCheck{}
}

func (h HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`{"ok":true}`))
}
