package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func user_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello User!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user", user_handler)
	http.ListenAndServe(":8080", nil)
}