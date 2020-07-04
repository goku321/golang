package main

import (
	"fmt"
	"strings"
)

func concatString(args ...interface{}) string {
	var b strings.Builder
	
	for _, v := range args {
		fmt.Fprintf(&b, "%v|", v)
	}

	return b.String()
}

func main() {
	str1 := concatString(1, 2, 3, "a", "b", "c", true)
	fmt.Println(str1)
}