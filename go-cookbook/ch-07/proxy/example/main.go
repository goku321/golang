package main

import "golang/go-cookbook/ch-07/proxy"

import "net/http"

import "fmt"

func main() {
	p := &proxy.Proxy{
		Client: http.DefaultClient,
		BaseURL: "https://www.golang.org",
	}
	http.Handle("/", p)
	fmt.Println("listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}