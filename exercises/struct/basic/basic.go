package main

import "log"

type jedi struct {
	name         string // name of the jedi.
	midichlorins int    // midichlorins count
	lightsaber   string // lightsaber color.
}

func main() {
	var anakin jedi
	// var definition assigns zero value?
	// we can access struct fields as well.
	log.Println(anakin, anakin.midichlorins)

	luke := jedi{}
	// struct comparison possible?
	if anakin == luke {
		log.Println("yes")
	}
}
