package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
	fmt.Printf("%d\t%q\t%t", x, y, z)
}