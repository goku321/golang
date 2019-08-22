package main

import "fmt"

func main() {
	x := [][]string{{"james", "bond"}, {"miss", "penny"}}

	for _, i := range x {
		for _, j := range i {
			fmt.Println(j)
		}
	}
}