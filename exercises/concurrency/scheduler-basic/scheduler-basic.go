// Program to show how the scheduler time slice goroutines
// on a single thread

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Start goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish...")
	wg.Wait()

	fmt.Println("Terminating main program")
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 4000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed ", prefix)
}
