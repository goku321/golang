package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	unbuffered := make(chan int)

	go func() {
		defer wg.Done()
		unbuffered <- 1
	}()

	go func() {
		defer wg.Done()
		value := <-unbuffered
		fmt.Println("Value received ", value)
	}()

	wg.Wait()
}
