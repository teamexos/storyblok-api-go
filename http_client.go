package storyblok

import (
	"net/http"
	"time"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// DefaultHTTPClient returns the default REST client
func DefaultHTTPClient() HTTPClient {
	return &http.Client{
		Timeout: time.Minute,
	}
}
