package handler

import (
	"net/http"

	"github.com/mhoc/urlshortener/pkg/store"
	"github.com/mhoc/urlshortener/pkg/util"
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
	w.Header().Add("Location", redirectTo)
	w.WriteHeader(302)
	w.Write([]byte(""))
}
