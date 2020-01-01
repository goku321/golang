package middleware

import "net/http"

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
