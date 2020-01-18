package proxy

import "net/http"

// Proxy holds configured client and BaseURL to proxy to
type Proxy struct {
	Client *http.Client
	BaseURL string
}