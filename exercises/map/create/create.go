package main

import "fmt"

func main() {
	// create a map using built-in function make.
	// also an empty map.
	m := make(map[string]int)
	m["luke"] = 25

	// create a map using map literal.
	n := map[string]int{
		"yoda": 100,
	}

	// remove a map element.
	delete(n, "yoda")

	// what if key is not present inside the map?
	// this will work - if key is not present a zero
	// value will be returned.
	m["han"] = m["han"] + 1
	fmt.Println(m)
	// this will work as well.
	m["ben"]++

	// compile error - not possible to take an address of map element.
	// _ := &m["ben"]

	// what's a zero value for a map?
	// let's see!
	var zeroMap map[string]int
	// It's nil - a zero map points to no hash table at all.
	fmt.Println(zeroMap == nil) // "true"
	// nil and zero length.
	fmt.Println(len(zeroMap) == 0) // "true"

	// one gotcha here - a nil map behaves like an empty map.
	// so all the operation except storing to a nil map are safe.
	// zeroMap["lea"] = 21 // will panic.
	// lesson learned - map vs var
	// map creates an empty map.
	// var creates a nil map.

	// what happens if the key is not present?
	// you will get the zero value of that type.
	fmt.Println(zeroMap["no"])

	// how would you know if the element is there or not?
	// use ok!
	_, ok := zeroMap["not"]
	if !ok {
		fmt.Printf("element with key '%s' is not present", "not")
	}

	// can you compare two map?
	// no. a map can only be compared to nil.
	// m == zeroMap // not possible.
}
