package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%T\n", x)
	fmt.Println(cap(x), len(x))

	for _, v := range x {
		fmt.Println(v)
	}
}
