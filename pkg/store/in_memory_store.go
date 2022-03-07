package store

import (
	"sync"
	"time"

	"gitlab.com/mhoc/urlshortener/pkg/util"
)

// InMemoryStore stores url redirects in-memory, expectedly losing them upon service restart.
//
// I create two maps here; one which maps short IDs to the redirected URL, and one for the oppsite.
// We do this in order to fulfil the requirement that two identical URLs should generate duplicate
// shortlinks.
//
// This store is also thread-safe, through the inclusion of a mutex on any creation/deletion
// operations.
type InMemoryStore struct {
	lock    sync.Mutex
	idToUrl map[string]string
	urlToId map[string]string
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		idToUrl: make(map[string]string),
		urlToId: make(map[string]string),
	}
}

func (ims *InMemoryStore) Create(redirectToUrl string, expiresIn time.Duration) (string, error) {
	ims.lock.Lock()
	defer ims.lock.Unlock()
	if existingId, in := ims.urlToId[redirectToUrl]; in {
		return existingId, nil
	}
	id := util.NewID()
	ims.idToUrl[id] = redirectToUrl
	ims.urlToId[redirectToUrl] = id
	return id, nil
}

func (ims *InMemoryStore) Get(id string) (string, error) {
	if redirectToUrl, in := ims.idToUrl[id]; in {
		return redirectToUrl, nil
	}
	return "", nil
}

func (ims *InMemoryStore) Remove(id string) (bool, error) {
	ims.lock.Lock()
	defer ims.lock.Unlock()
	if redirectToUrl, in := ims.idToUrl[id]; in {
		delete(ims.idToUrl, id)
		delete(ims.urlToId, redirectToUrl)
		return true, nil
	}
	return false, nil
}

func (ims *InMemoryStore) Stop() {
}
