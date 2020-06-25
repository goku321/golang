package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	errNoRows = errors.New("db: no rows found")
	errDupRow = errors.New("db: duplicate row")
)

type dbError struct {
	method string
	err    error
}

func (e *dbError) Error() string {
	return fmt.Sprintf("%s\n", e.method, e.err)
}

func getUsers() *dbError {
	return &dbError{
		method: "getUsers",
		err:    errNoRows,
	}
}

func saveUser() error {
	return fmt.Errorf("saveUser: %w", errDupRow)
}

func main() {
	err := getUsers()
	if _, ok := err.(*dbError); !ok {
		log.Print("yeah")
	}
}
