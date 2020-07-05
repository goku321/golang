package main

import "fmt"

type x interface {
	get() string
	put(string)
}

type z struct {}

func (i z) put(s string) {
	fmt.Println("putting: ", s)
}

func (i z) get() string {
	return "nay"
}

type y struct {
	x
}

func (i y) get() string {
	return "yay"
}

func (i y) put(s string) {
	fmt.Println("put put")
}

func main() {
	var a x // is of type interface.
	var c z
	b := y{x: c} // is a struct.
	a = b

	fmt.Println(a.get())
	a.put("value")
	fmt.Println(b.get())
	fmt.Println(b.x.get())
	b.put("")
	b.x.put("")
}
