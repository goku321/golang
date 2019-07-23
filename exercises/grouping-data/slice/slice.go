package main

import "fmt"

func main() {
	// Creating a slice using a COMPOSITE LITERAL
	x := []int{1, 3, 4, 5}
	
	// Iterating over a slice using range clause
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Slicing a slice using colon operator
	slicedSlice := x[1:3]

	fmt.Println(slicedSlice)
}