package proxy

import (
	"bytes"
	"net/http"
	"net/url"
)

// ProcessRequest modifies the request in accordance with the proxy settings
func (p *Proxy) ProcessRequest(r *http.Request) error {
	proxyURLRaw := p.BaseURL + r.URL.String()

	proxyURL, err := url.Parse(proxyURLRaw)
	if err != nil {
		return err
	}
	r.URL = proxyURL
	r.Host = proxyURL.Host
	r.RequestURI = ""

	return nil
}

// CopyResponse takes the client response and writes everything to the
// ResponseWriter in the original handler
func CopyResponse(w http.ResponseWriter, resp *http.Response) {
	var out bytes.Buffer
	out.ReadFrom(resp.Body)

	for k, v := range resp.Header {
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(out.Bytes())
}
