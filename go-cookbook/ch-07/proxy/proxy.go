package proxy

import (
	"log"
	"net/http"
)

// Proxy holds configured client and BaseURL to proxy to
type Proxy struct {
	Client  *http.Client
	BaseURL string
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := p.ProcessRequest(r); err != nil {
		log.Printf("error occurred during processing request %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
