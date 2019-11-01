/* A hello world web app in Go
   Two handlers are defined for two different endpoints.
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/uptime"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func userHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello User!")
}

func main() {
	h := health.New(health.Health{Version: "1", ReleaseID: "1.0.0"}, uptime.System())
	http.HandleFunc("/", handler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/health", h.Handler)
	http.ListenAndServe(":8080", nil)
}