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

func main() {
	tail := &intList{4, nil}
	head := &intList{2, tail}
	fmt.Println(head.sum())
}
