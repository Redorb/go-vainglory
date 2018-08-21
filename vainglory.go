package vainglory

import (
	"net/http"
	"net/url"
)

// Client is the main struct for vainglory
type Client struct {
	apiKey     string // developers api key, used to make calls
	httpClient *http.Client
	baseURL    *url.URL
}

// New returns a new defaulted Session struct.
func New(key string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	url, err := url.Parse("https://api.dc01.gamelockerapp.com/")
	if err != nil {
		return nil, err
	}

	return &Client{
		apiKey:     key,
		httpClient: httpClient,
		baseURL:    url,
	}, nil
}
