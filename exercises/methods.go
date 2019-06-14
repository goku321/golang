package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) getName() string {
	return p.name
}

func (p person) getAge() int {
	return p.age
}

func main() {
	s := person{"Deepak", 25}
	fmt.Println(s.getName())
	fmt.Println(s.getAge())
}
