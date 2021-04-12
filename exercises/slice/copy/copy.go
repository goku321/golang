// This example shows how to use copy built-in
// to create a room in a slice to insert a element
// at a given index.
package main

import "fmt"

// The slice must have room for the new element.
func insert(x []int, index, value int) []int {
	// Grow the slice by one element.
	x = x[0 : len(x)+1]
	// Create room for new element.
	copy(x[index+1:], x[index:])
	// Insert the new element.
	x[index] = value
	return x
}

func main() {
	x := make([]int, 5, 10)
	for i := range x {
		x[i] = i
	}
	fmt.Println(insert(x, 2, 3))
}
