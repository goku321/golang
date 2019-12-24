package main

import (
	"fmt"
	"golang/go-cookbook/ch-07/validation"
	"net/http"
)

func main() {
	c := validation.New()
	http.HandleFunc("/", c.Process)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
