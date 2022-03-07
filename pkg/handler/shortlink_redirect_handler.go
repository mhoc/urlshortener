package handler

import "net/http"

type ShortlinkRedirect struct{}

func NewShortlinkRedirect() ShortlinkRedirect {
	return ShortlinkRedirect{}
}

func (h ShortlinkRedirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
