// This example shows how to use copy built-in
// to create a room in a slice to insert a element
// at a given index.
package main

import "fmt"

// The slice must have room for the new element.
func insert(x []int, index, value int) []int {
	// Grow the slice by one element.
	x = x[0 : len(x)+1]
	copy(x[index+1:], x[index:])
	x[index] = value
	return x
}

func main() {
	x := [5]int{1, 2, 4, 5}
	fmt.Println(insert(x[:4], 2, 3))
}
