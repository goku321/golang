package main

import (
	"fmt"
	"net/http"
)

// SuccessHandler is called on success.
type SuccessHandler struct{}

func (s *SuccessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request Success")
}

// FailureHandler is called on failure.
type FailureHandler struct{}

func (f *FailureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request Failure")
}

func main() {
	s := SuccessHandler{}
	f := FailureHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/success", &s)
	http.Handle("/failure", &f)

	server.ListenAndServe()
}
