package storyblok

import (
	"net"
	"net/http"
	"time"
)

// HTTP client defaults
var (
	defaultHTTPTransportDisableCompression = false
	defaultHTTPTransportIdleConnTimeout    = 30 * time.Second
	defaultHTTPTransportMaxIdleConns       = 5
	defaultHTTPTransportDialContextTimeout = 5 * time.Second
	defaultHTTPClientTimeout               = time.Minute
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// DefaultHTTPClient returns the default REST client
func DefaultHTTPClient() HTTPClient {

	// Instantiate gzip client with a 5 second timeout on waiting for the
	// remote server to accept the connection and a 30 second timeout
	// for no activity over the connection
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: defaultHTTPTransportDialContextTimeout,
		}).DialContext,
		MaxIdleConns:       defaultHTTPTransportMaxIdleConns,
		IdleConnTimeout:    defaultHTTPTransportIdleConnTimeout,
		DisableCompression: defaultHTTPTransportDisableCompression,
	}

	return &http.Client{
		Timeout:   defaultHTTPClientTimeout,
		Transport: transport,
	}
}
