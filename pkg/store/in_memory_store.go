package store

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/mhoc/urlshortener/pkg/util"
)

// InMemoryStore stores url redirects in-memory, expectedly losing them upon service restart.
//
// I create two maps here; one which maps short IDs to the redirected URL, and one for the oppsite.
// We do this in order to fulfil the requirement that two identical URLs should generate duplicate
// shortlinks.
//
// All that is to say; this is a pretty memory-heavy store. Its a trade-off; some of these
// structures could be removed; less memory but higher latencies on certain operations.
//
// This store is also thread-safe, through the inclusion of a mutex on any creation/deletion
// operations.
type InMemoryStore struct {
	lock    sync.Mutex
	idToUrl map[string]string
	urlToId map[string]string
}

func NewInMemoryStore() *InMemoryStore {
	ims := &InMemoryStore{
		idToUrl: make(map[string]string),
		urlToId: make(map[string]string),
	}
	return ims
}

func (ims *InMemoryStore) Create(ctx context.Context, redirectToUrl string, expiresIn time.Duration) (string, error) {
	// If the redirect url already exists in our store, we simply return the existing id.
	if existingId, in := ims.urlToId[redirectToUrl]; in {
		return existingId, nil
	}
	// Grab the mutex lock, since we're doing a write operation.
	ims.lock.Lock()
	defer ims.lock.Unlock()
	// Generate a new ID for the redirect url. This will form the entire path component of the final
	// shortened url.
	id := util.NewID()
	// Map that ID to the URL we want to redirect to, and conversely map the URL back to the ID for
	// quicker deletion.
	ims.idToUrl[id] = redirectToUrl
	ims.urlToId[redirectToUrl] = id
	// ExpiresIn will be positive if we desire this shortened URL to expire at some point in the
	// future.
	if expiresIn > 0 {
		// We accomplish this as simply as possible here, via time.AfterFunc.
		time.AfterFunc(expiresIn, func() {
			// We don't need to grab the mutex during the removal, as ims.Remove() will grab it
			// anyway. The error handling here isn't the best.
			_, err := ims.Remove(context.Background(), id)
			if err != nil {
				log.Printf("error expiring %v: %v", id, err.Error())
			}
		})
	}
	return id, nil
}

func (ims *InMemoryStore) Get(ctx context.Context, id string) (string, error) {
	// Fairly straightforward; if the ID exists in our mapping of IDs to URLs, we can return it, but
	// otherwise we want to return an empty string.
	if redirectToUrl, in := ims.idToUrl[id]; in {
		return redirectToUrl, nil
	}
	return "", nil
}

func (ims *InMemoryStore) Remove(ctx context.Context, id string) (bool, error) {
	// A write operation, and thus we want to grab the mutex.
	ims.lock.Lock()
	defer ims.lock.Unlock()
	// The easy case here is where we are removing an entry that doesn't exist in the in-memory
	// store. We simply return false and call it a day.
	if _, in := ims.idToUrl[id]; !in {
		return false, nil
	}
	redirectToUrl := ims.idToUrl[id]
	// However, if it does exist, we need to remove the id from every "index" we've created within
	// the in-memory store. So: The index which maps IDs to URLs.
	delete(ims.idToUrl, id)
	// and the index which maps URLs to IDs.
	delete(ims.urlToId, redirectToUrl)
	return true, nil
}

func (ims *InMemoryStore) Stop() {
	// This currently does nothing.
	// Theoretically, this should be improved to store the list of timers created above to handle
	// expiration, then stop all of them. But, considering this only exists in order to write tests
	// against and be useful during local dev, im classifying it as "save for v2".
}
