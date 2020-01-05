package main

import (
	"fmt"
	"golang/go-cookbook/ch-07/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	h := middleware.ApplyMiddleware(
		middleware.Handler,
		middleware.Logger(log.New(os.Stdout, "", 0)),
		middleware.SetID(100),
	)

	http.HandleFunc("/", h)
	fmt.Println("listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
