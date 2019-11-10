package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is: ", v.Kind())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)

	type MyInt int
	var y MyInt = 7
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.ValueOf(y).Kind())
}