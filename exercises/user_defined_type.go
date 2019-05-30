package main

import (
	"fmt"
)

type dragon int
var water int
var fire dragon

func main() {
	fmt.Println(fire)
	fmt.Printf("%T\n", fire)
	fire = 42
	fmt.Printf("%v\n", fire)
	water = int(fire)
	fmt.Println("water: ", water)
}