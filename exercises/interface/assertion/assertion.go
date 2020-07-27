package main

import (
	"log"
	"reflect"
)

func main() {
	var v interface{}
	v = "4"

	var val int
	var ok bool
	if val, ok = v.(int); !ok {
		log.Fatalf("cannot assert value of type %s to an int value", reflect.TypeOf(v))
	}
	log.Println(val)
}
