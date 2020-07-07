package main

import "fmt"

type sand struct {
	character []int
}

func main() {
	m := map[string]sand{}

	m["clay"] = sand{
		character: []int{1},
	}

	// update character.
	v := m["clay"]
	v.character = append(v.character, 5)
	fmt.Println(m)

	m["clay"] = v
	fmt.Println(m)
}
