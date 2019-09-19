package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(concat("a", "b", "c"))
	fmt.Println(concat())
}

func concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied")
	}

	return strings.Join(parts, " "), nil
}
