package client

import (
	"crypto/tls"
	"net/http"
)

// Setup configures our client and redefines
// the global DefaultClient
func Setup(isSecure, nop bool) *http.Client {
	c := http.DefaultClient

	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}

	if nop {
		c.Transport = &NopTransport{}
	}

	http.DefaultClient = c
	return c
}

// NopTransport is a No-Op Transport
type NopTransport struct{}

// RoundTrip implements RoundTripper interface
func (n *NopTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}
