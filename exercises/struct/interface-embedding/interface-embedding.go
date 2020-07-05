package main

import "fmt"

type x interface {
	get() string
	put(string)
}

type y struct {
	x
}

// func (i y) get() string {
// 	return "yay"
// }

// func (i y) put(s string) {
// 	fmt.Println("put put")
// }

type z struct{}

func (i z) get() string {
	return "from z"
}

func (i z) put(s string) {
	fmt.Println("from z: ", s)
}

func main() {
	var a x   // x is an interface.
	var c z   // z implements x.
	b := y{c} // injecting z for x.

	a = b // b is of type y and y embeds x, so it automatically implements x.

	fmt.Println(a.get()) // will call get method on z.
	a.put("value")       // will call put method on z.

	fmt.Println(b.x.get()) // will call get method on z.
	b.x.put("value")       // will call put method on z.
	fmt.Println(b.get())   // since y doesn't explicitly defined get method, it fallbacks to z.get().
	b.put("value")
}
