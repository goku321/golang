package main

import "fmt"

type vehicle struct {
	color string
}

type car struct {
	vehicle
	color string
}

type bike struct {
	vehicle
	wheels int
}

func main() {
	dalorean := car{
		vehicle{
			color: "grey",
		},
		"blue",
	}
	fmt.Println(dalorean.color, dalorean, dalorean.vehicle, dalorean.vehicle.color)

	ford := car {
		vehicle{
			"grey",
		},
		"yellow",
	}
	fmt.Println(ford.color, ford, ford.vehicle, ford.vehicle.color)

	ducati := bike{
		vehicle{
			"grey",
		},
		2,
	}
	fmt.Println(ducati.color, ducati.vehicle.color, ducati.vehicle, ducati.wheels)
}
