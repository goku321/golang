package main

import (
	"fmt"
)

func main() {
	unbuffered := make(chan int)
	
	go func() {
		unbuffered <- 1
	}()

	go func() {
		value := <-unbuffered
		fmt.Println("Value received from channel: ", value)
	}()
}
