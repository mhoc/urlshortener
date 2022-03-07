package util

import (
	"math/rand"
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
