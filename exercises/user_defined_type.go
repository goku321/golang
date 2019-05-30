package main

import (
	"fmt"
)

type dragon int

func main() {
	var fire dragon
	fmt.Println(fire)
	fmt.Printf("%T\n", fire)
	fire = 42
	fmt.Printf("%v\n", fire)
}