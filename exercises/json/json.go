package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type person struct {
	Age int
}

func main() {
	var p *person
	fmt.Println(json.Unmarshal([]byte("null"), p).Error())
	fmt.Println(p, &p)

	fmt.Println(ioutil.ReadFile(""))
}
