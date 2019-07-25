package main

import "fmt"

func main() {
	// Creating a slice using a COMPOSITE LITERAL
	x := []int{1, 3, 4, 5}

	// Iterating over a slice using range clause
	fmt.Printf("Iterating over slice: %v\n", x)
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Slicing a slice using colon operator
	slicedSlice := x[1:3]

	fmt.Printf("Sliced slice: %v\n", slicedSlice)

	// Appending to a slice using built-in `append`
	x = append(x, 6, 7, 8, 9)
	fmt.Printf("Slice x after appending: %v\n", x)

	// Appending one slice into another
	y := []int{99, 100, 101}
	x = append(x, y...)
	fmt.Printf("Slice x after appending y to it: %v\n", x)
}
