package main

import "fmt"

func main() {
	expr := "favSport"

	switch expr {
	case "favSport":
		fmt.Println("Should Print")
	case "favColor":
		fmt.Println("Should Not Print")
	}
}
