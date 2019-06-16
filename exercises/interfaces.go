package main

import "fmt"

/* Any type that defines `Speak` method
is said to satisfy the Animal interface */
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Cow struct {}

func (c Cow) Speak() string {
	return "Mooo..."
}

func main() {
	// animals := []Animal{Dog{}, Cat{}} won't work
	animals := []Animal{Dog{}, new(Cat), new(Cow)}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	d := &Dog{}
	fmt.Printf("%T\n", d)

	c := &Cat{}
	fmt.Println(c.Speak())
}
