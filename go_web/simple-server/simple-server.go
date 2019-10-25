/* A hello world web app in Go
   Two handlers are defined for two different endpoints.
*/

package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func userHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello User!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}