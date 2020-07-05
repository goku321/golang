package main

import "fmt"

/* Any type that defines `Speak` method
is said to satisfy the Animal interface */
type animal interface {
	speak() string
}

type dog struct{}

func (d dog) speak() string {
	return "Woof!"
}

type cat struct{}

func (c *cat) speak() string {
	return "Meow!"
}

type cow struct {}

func (c cow) speak() string {
	return "Mooo..."
}

func (c cow) milk() string {
	return "here's my milk"
}

func main() {
	// animals := []animal{Dog{}, Cat{}} won't work
	animals := []animal{dog{}, &cat{}, new(cow)}

	for _, animal := range animals {
		fmt.Println(animal.speak())
	}

	d := &dog{}
	fmt.Printf("%T\n", d)

	c := &cat{}
	fmt.Println(c.speak())

	moo := cow{}
	fmt.Println(moo.milk())
}
