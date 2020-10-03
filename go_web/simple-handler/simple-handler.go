// Implementing a HTTP Handler

package main

import (
	"fmt"
	"net/http"
	"os"
)

// MyHandler is an instance of Handler since it has a ServeHTTP method
type MyHandler struct {
	requestCount int
}

func (h *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("received request from %s\n", request.RemoteAddr)
	if h.requestCount == 5 {
		writer.WriteHeader(500)
		writer.Write([]byte("internal server error"))
		return
	}
	host, err := os.Hostname()
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(200)

	writer.Write([]byte(fmt.Sprintf("hello world from %s", host)))
	h.requestCount++
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
