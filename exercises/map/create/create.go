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
	fmt.Println(zeroMap == nil)
}
