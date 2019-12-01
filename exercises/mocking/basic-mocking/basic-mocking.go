package main

import "fmt"

type Stringer interface {
	String() string
	PrintString(string)
}

type s string

func (a s) String() string {
	return "test string"
}

func (a s) PrintString(x string) {
	fmt.Println(x)
}

func GetString(x Stringer) string {
	return x.String()
}

func Display(x Stringer) {
	x.PrintString("hi there")
}

func main() {
	var y s
	GetString(y)
	Display(y)
}
