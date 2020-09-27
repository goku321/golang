package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	defer resp.Body.Close()
	buf.ReadFrom(resp.Body)
	downloadedBytes := buf.Bytes()
	fmt.Println(downloadedBytes)
	fmt.Println(buf.String())

}
