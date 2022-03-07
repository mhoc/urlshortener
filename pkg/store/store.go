package store

import "time"

// Store is an interface which the data storage mechanisms within the url shortener fulfill.
//
// I've created two different stores at this time; a Redis-backed one which will persist during
// service and system restarts (assuming redis is configured properly of course), and an in-memory
// one which is useful for testing and local development.
type Store interface {
	// Create should create a new shortlink within the store which redirects to the provided URL,
	// and return back the ID of the provided mapping, or an error if applicable.
	//
	// Create should not create new entries for destination urls which already exist in the store,
	// and should instead return the existing entry.
	//
	// expiresIn can be provided to instruct the store to purge the shortlink after the given
	// duration has passed. If expiresIn is provided to `0`, the shortlink does not expire.
	Create(redirectToUrl string, expiresIn time.Duration) (string, error)

	// Get should retrieve the requested shortlink by ID, and return back the URL it redirects to.
	//
	// If the requested ID does not correspond to a shortlink, including if the shortlink it
	// corresponds to has expired, an empty string and no error is returned.
	Get(id string) (string, error)

	// Remove should remove the requested shortlink by id, returning `true` if the entry existed
	// in the store and was removed, or false and no error if the entry did not exist in the store
	// at all.
	Remove(id string) (bool, error)

	// Stop should shut down any connections to the store, including
	Stop()
}
