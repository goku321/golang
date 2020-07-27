package main

import (
	"log"
	"reflect"
)

func main() {
	var v interface{}
	v = true

	var val int
	var ok bool
	if val, ok = v.(int); !ok {
		log.Printf("cannot assert value of type %s to an int value", reflect.TypeOf(v))
	}
	log.Println(val)

	// Switch on interface value.
	switch v {
	case "4":
		log.Println("switch works")
	case true:
		log.Println("works on booleans as well")
	}

	// let's try asserting an int value to float64.
	v = 5
	var fv float64
	if fv, ok = v.(float64); !ok {
		log.Printf("cannot assert value of type %s to float64 value", reflect.TypeOf(v))
	}
	log.Println(fv)
}
