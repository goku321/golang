package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
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

	// Embedded struct
	sa1 := secretAgent{
		person: person{
			first: "Deepak",
			last:  "Sah",
			age:   25,
		},
		ltk: false,
	}

	fmt.Println("SA1: ", sa1)
	fmt.Println	(sa1.first, sa1.last, sa1.ltk)
}
