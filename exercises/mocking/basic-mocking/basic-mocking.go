package main

import "fmt"

type Stringer interface {
	String() string
}

func main() {
	var a s
	fmt.Println(GetString(a))
}

type s string

func (a *s) String() string {
	return "test string"
}

func GetString(x s) string {
	return x.String()
}