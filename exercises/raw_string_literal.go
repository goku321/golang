package main

import (
	"fmt"
)

func main() {
	var x string = `"A string literal"
	which can expand to
	multiple lines`
	fmt.Println(x)
}