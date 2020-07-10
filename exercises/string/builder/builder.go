package main

import (
	"fmt"
	"strings"
	"time"
)

func concatString(args ...interface{}) string {
	var b strings.Builder

	for _, v := range args {
		fmt.Fprintf(&b, "%v|", v)
	}

	return b.String()
}

func main() {
	str1 := concatString(1, string(2), 3, "a", "b", "c", true)
	fmt.Println(str1)

	str2 := concatString(time.Now(), 1, 2, 3, "a", true)
	fmt.Println(str2)
}
