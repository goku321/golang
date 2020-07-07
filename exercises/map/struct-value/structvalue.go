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

	// does not update the map.
	v := m["clay"]
	v.character = append(v.character, 5)
	fmt.Println(m)

	// assign the whole value to update.
	m["clay"] = v
	fmt.Println(m)
}
