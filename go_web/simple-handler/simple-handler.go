// Implementing a HTTP Handler

package main

import (
	"fmt"
	"net/http"
	"os"
)

// MyHandler is an instance of Handler since it has a ServeHTTP method
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("received request from %s\n", request.RemoteAddr)
	host, err := os.Hostname()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(200)

	writer.Write([]byte(fmt.Sprintf("hello world from %s", host)))
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: &handler,
	}
	fmt.Println("listening on :8081")
	server.ListenAndServe()
}
