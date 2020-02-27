package main

import "golang/go-cookbook/ch-06/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
