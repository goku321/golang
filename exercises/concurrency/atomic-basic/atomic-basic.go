package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown now...")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

func doWork(id string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", id)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", id)
			break
		}
	}
}
