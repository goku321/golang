package main

import "fmt"

func main() {
	for i := 33; i < 123; i++ {
		fmt.Printf("%q\t%#U\n", i, i)
	}
}
