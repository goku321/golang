package main

import (
	"fmt"
)

const (
	a  int = 42
	b string = "hello"
	c float64 = 42.42
)

func main() {
	fmt.Printf("%d\n", a)
	fmt.Printf("%s\n", b)
	fmt.Printf("%f\n", c)
}