package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter will be incremented by multiple goroutines
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		// Yield the thread and be placed back in queue.
		runtime.Gosched()

		value++

		counter = value
	}
}
