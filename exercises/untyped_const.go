package main

import (
	"fmt"
)

const (
	a = 42
	b = "hello"
	c = 42.42
)

func main() {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}