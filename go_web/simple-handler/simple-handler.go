// Implementing a HTTP Handler

package main

import (
	"fmt"
	"net/http"
)

// MyHandler is an instance of Handler since it has a ServeHTTP method
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hey")
	writer.WriteHeader(200)

	writer.Write([]byte("hello world"))
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: &handler,
	}
	server.ListenAndServe()
}
