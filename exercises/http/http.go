package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("failed to download sampling strategies: %v", err)
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		fmt.Printf("failed to read sampling strategies HTTP response body: %v", err)
	}
	fmt.Println(buf.String())
}
