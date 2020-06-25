package main

import (
	"errors"
	"fmt"
)

type dbError struct {
	query string
	err   error
}

func (e *dbError) Error() string {
	return fmt.Sprintf("error while executing query %s:\n%s", e.query, e.err.Error())
}

func main() {
	var d *dbError
	d = &dbError{
		query: "select * from table",
		err: errors.New("some serious error"),
	}
	fmt.Println(d.Error())
}
