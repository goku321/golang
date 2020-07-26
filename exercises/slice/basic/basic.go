package main

import "log"

func main() {
	var x []int
	// what is zero value of slice when defined using var?
	log.Println(x)
	// what about length and capacity?
	log.Printf("len: %v, cap: %v", len(x), cap(x))
	// how about getting a value using an index?
	// log.Println(x[0]) // panics

	// And when defined like this?
	y := []int{}
	log.Println(y)
	log.Printf("len: %v, cap: %v", len(y), cap(y))
	log.Println(y == nil)
	// let's try with y.
	// log.Println(y[0]) // panics

	// can nil be assigned to a slice?
	x = nil
	log.Println(x)
	log.Printf("len: %v, cap: %v", len(x), cap(x))
	log.Println(x == nil)
}