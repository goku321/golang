package main

import (
	"fmt"
	"golang/go-cookbook/ch-07/negotiate"
	"net/http"
)

func main() {
	http.HandleFunc("/", negotiate.Handler)
	fmt.Println("listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
