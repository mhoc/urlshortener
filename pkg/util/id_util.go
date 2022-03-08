package util

import (
	"fmt"
	"math/rand"
	"net/url"

	"time"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
)

var (
	prng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// NewID generates a new random identifier for a url. Its hardcoded at this time to only generate
// 8 character IDs, which is pretty short. Given the current alphabet size of 64 characters, that's
// a couple hundred trillion ids I think.
func NewID() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = alphabet[prng.Intn(len(alphabet))]
	}
	return string(b)
}

// IDToShortlink converts an ID, and the root url of the server, into a shortlink viable to returned
// to the client for url shortening fun.
func IDToShortlink(rootUrl *url.URL, id string) string {
	return rootUrl.String() + "/" + id
}

// ShortlinkToID takes a shortlink and retrieves the ID from the path
func ShortlinkToID(shortlink string) (string, error) {
	u, err := url.Parse(shortlink)
	if err != nil {
		return "", fmt.Errorf("unable to parse url shortlink: %v", err.Error())
	}
	// u.Path includes the leading slash
	return u.Path[1:], nil
}
