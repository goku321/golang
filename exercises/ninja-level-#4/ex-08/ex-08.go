package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`: []string{"shaken", "not stirred", "martini"},
		`moneypenny_miss`: []string{"James Bond", "Literature", "Computer Science"},
	}

	for k, v := range x {
		fmt.Printf("%v: ", k)
		for _, i := range v {
			fmt.Printf("%v ", i)
		}
		fmt.Println()
	}
}