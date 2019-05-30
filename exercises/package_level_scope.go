package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println(x, y, z)
	fmt.Printf("%d\t%q\t%t", x, y, z)
	s := fmt.Sprintf("%d\t%q\t%t", x, y, z)
	fmt.Println(s)
}