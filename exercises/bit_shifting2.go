package main

import (
	"fmt"
)

func main() {
	x := 49
	fmt.Printf("%d\t%b\t%x\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\t%b\t%x", y, y, y)
}