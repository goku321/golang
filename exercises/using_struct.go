package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Deepak", 20})
	fmt.Println(person{name: "Joe", age: 22})
	fmt.Println(person{age: 30})
	fmt.Println(&person{"Mike", 33})
	s := person{"Brad", 34}
	fmt.Println(s.name)
	s.age = 42
	fmt.Println(s.age)

}
