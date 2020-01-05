package middleware

import (
	"log"
	"time"
	"net/http"
)

// Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ApplyMiddleware will apply all the middlewares
func ApplyMiddleware(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	applied := h
	for _, m := range middleware {
		applied = m(applied)
	}
	return applied
}

// Logger logs requests
func Logger(l *log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Printf("started request to %s with id %s", r.URL, GetID(r.Context()))
			next(w, r)
			l.Printf("completed request to %s with id %s in %s", r.URL, GetID(r.Context()), time.Since(start))
		}
	}
}
