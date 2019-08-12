package main

import "fmt"

func main() {
	// Creating a string to int map
	m := map[string]int{
		"Deepak": 6,
		"Sah":    3,
	}

	// Printing a map
	fmt.Printf("Map m looks like: %v\n", m)

	// Accessing a map using a key
	fmt.Println(m["Deepak"])

	// Accessing a non-existing key from a map
	fmt.Println(m["Milton"])

	// Indexing a map returns two values
	value, ok := m["Milton"]
	fmt.Println(value, ok)
	nv, ok := m["Sah"]
	fmt.Println(nv, ok)

	if nv, ok := m["Deepak"]; ok {
		fmt.Printf("Age of Deepak: %v\n", nv)
	}

	// Adding an element to a map
	m["Todd"] = 33
	fmt.Printf("Map after adding an element: %v\n", m)

	// Iterating over a map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Deleting from a map
	delete(m, "Deepak")
	fmt.Printf("Map m after deleting key Deepak %v\n", m);

	// Deleting a non-existent key
	delete(m, "Tony")

	// Verifying that a key/value exist before deleting
	if value, ok := m["Sah"]; ok {
		fmt.Printf("Deleting value %v\n", value)
		delete(m, "Tony")
	}
}
