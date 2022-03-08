package handler

import (
	"net/http"

	"github.com/mhoc/urlshortener/pkg/store"
	"github.com/mhoc/urlshortener/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ShortlinkRedirectCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "api",
		Name:      "shortlink_redirects",
	}, []string{
		"short_url",
	})
)

type ShortlinkRedirect struct {
	st store.Store
}

func NewShortlinkRedirect(st store.Store) ShortlinkRedirect {
	return ShortlinkRedirect{
		st: st,
	}
}

func (h ShortlinkRedirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := util.ShortlinkToID(r.URL.String())
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"provided a malformed shortlink"}`))
		return
	}
	redirectTo, err := h.st.Get(r.Context(), id)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"error":"internal server error"}`))
		return
	}
	if redirectTo == "" {
		w.WriteHeader(404)
		w.Write([]byte(`{"error": "shortlink not found"}`))
		return
	}
	ShortlinkRedirectCounter.With(prometheus.Labels{"short_url": r.URL.String()}).Inc()
	w.Header().Add("Location", redirectTo)
	w.WriteHeader(302)
	w.Write([]byte(""))
}
