package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Deepak",
		last:  "Sah",
		age:   25,
	}

	fmt.Println("Person struct: ", p1)

	// Displaying by attribute
	fmt.Println(p1.first, " ", p1.last, " is ", p1.age, " years old")
}
