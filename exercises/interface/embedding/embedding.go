package main

type a interface {
	get() string
}

// cannot define another get method as method
// overloading is not allowed in go.
type b interface {
	a
	put(string)
}

type impl struct {}

func (i impl) put(s string) {
}

func (i impl) get() string {
	return ""
}

func main() {
	var x impl
	var y b
	y = x

	y.get()
	y.put("")
}