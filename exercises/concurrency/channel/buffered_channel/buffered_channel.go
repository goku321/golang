package main

import (
	"fmt"
)

func main() {
	buffered := make(chan string, 10)
	
	go func() {
		buffered <- "Gopher"
	}()

	go func() {
		value := <-buffered
		fmt.Println("Value received from channel: ", value)
	}()
}
