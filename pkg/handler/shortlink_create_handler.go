package handler

import "net/http"

type ShortlinkCreate struct{}

func NewShortlinkCreate() ShortlinkCreate {
	return ShortlinkCreate{}
}

func (h ShortlinkCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
