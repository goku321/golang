package main

import "fmt"

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func main() {
	m := map[string]int{}

	key1 := []string{"hi", "bye"}
	m[k(key1)] = 1
	fmt.Println(m[k(key1)])

	key2 := []string{"bye", "hi"}
	m[k(key2)] = 4
	fmt.Println(m[k(key2)])

	m[k(key1)] = 5
	fmt.Println(m)
}
