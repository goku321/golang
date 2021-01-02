package main

import (
	"golang/exercises/scope/user"
)

func main() {
	c := user.NewClient()
	user.NewReporter(c)
	t := &user.User{}
	t.Name("deepak")
}
