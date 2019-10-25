// Implementing a HTTP Handler

package main

import (
	"fmt"
	"net/http"
)

// MyHandler is an instance of Handler since it has a ServeHTTP method
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "I am being Served!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
