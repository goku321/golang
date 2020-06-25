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
	return fmt.Sprintf("%s\n%s", e.method, e.err)
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
	var dbErr *dbError
	if errors.As(err, &dbErr) {
		log.Println("it is db Error")
	}
	if errors.Is(err, errNoRows) {
		log.Printf("it is %v:", errNoRows)
	}
}
