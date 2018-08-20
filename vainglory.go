package vainglory

// Quite a bit of the code is borrowed from: https://github.com/LtSnuggie/pubgo
// Both PUBG and Vainglory use the same provider for their public apis.

import (
	"net/http"
	"time"
)

// Session is the main struct for pubgo
type Session struct {
	apiKey string  // developers api key, used to make calls
	poller *poller // poller is responsible for executing the requests as well as maintaining rate limit queue
}

// New returns a new defaulted Session struct.
func New(key string, rateLimit int) (s *Session, err error) {
	s = &Session{
		apiKey: key,
		poller: newPoller(&http.Client{Timeout: (20 * time.Second)}, rateLimit),
	}
	return
}
