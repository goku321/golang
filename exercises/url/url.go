package main

import (
	"fmt"
	"net/url"
)

func isURL(str string) bool {
    u, err := url.Parse(str)
    return err == nil && u.Scheme != "" && u.Host != ""
}

func main() {
	fmt.Println(url.Parse("randomstring"))
	fmt.Println(url.ParseRequestURI("/foo/bar"))
	fmt.Println(isURL("randomstring"))
	fmt.Println(isURL("/foo/bar"))
}
