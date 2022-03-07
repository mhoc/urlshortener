package handler

import "net/http"

type ShortlinkRemove struct{}

func NewShortlinkRemove() ShortlinkRemove {
	return ShortlinkRemove{}
}

func (h ShortlinkRemove) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
