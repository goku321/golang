package negotiate

import (
	"net/http"

	"github.com/unrolled/render"
)

// Negotiator ...
type Negotiator struct {
	ContentType string
	*render.Render
}

// GetNegotiator ...
func GetNegotiator(r *http.Request) *Negotiator {
	contentType := r.Header.Get("Content-Type")

	return &Negotiator{
		ContentType: contentType,
		Render: render.New(),
	}
}
