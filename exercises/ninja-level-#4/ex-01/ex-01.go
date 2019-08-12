package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("%v %v\n", i, v)
	}

	fmt.Printf("%T\n", a)
}
