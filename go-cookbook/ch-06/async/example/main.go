package main

import (
	"fmt"
	"golang/go-cookbook/ch-06/async"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://golang.org",
		"https://www.github.com",
	}

	c := async.NewClient(http.DefaultClient, len(urls))
	async.FetchAll(urls, c)

	for i := 0; i < len(urls); i++ {
		select {
		case resp := <-c.Resp:
			fmt.Printf("Status received for %s: %d\n",
				resp.Request.URL, resp.StatusCode)
		case err := <-c.Err:
			fmt.Printf("Error received: %s\n", err)
		}
	}
}
