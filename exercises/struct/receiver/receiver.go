package main

import "fmt"

type intList struct {
	value int
	tail  *intList
}

// Allowing nil as a valid receiver value.
func (list *intList) sum() int {
	if list == nil {
		return 0
	}
	return list.value + list.tail.sum()
}

type values map[string][]string

func (v values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	tail := &intList{4, nil}
	head := &intList{2, tail}
	fmt.Println(head.sum())

	// using values type
	m := values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang"))
	// direct access.
	fmt.Println(m["item"])

	m = nil
	// m.Add("item", "3") // Panics as entry in nil map is not allowed.
	fmt.Println(len(m))
}
