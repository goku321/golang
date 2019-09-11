package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Allocate one logical processor
	runtime.GOMAXPROCS(2)

	// Use wg to wait for the program to finish
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Goroutine #1
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Goroutine #2
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Println("Waiting to finish...")
	}

	wg.Wait()

	fmt.Println("\n Terminating program")
}
