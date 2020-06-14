package main

import "time"

// Person represents a person.
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
	BornAt    time.Time
}

// JSONPerson represents marshalled Person.
type JSONPerson struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
	BornAt    int64
}

func (jd JSONPerson) person() Person {
	return Person{
		jd.FirstName,
		jd.LastName,
		jd.Age,
		jd.Gender,
		time.Unix(jd.BornAt, 0),
	}
}

// MarshalJSON marshalls Person.
func (p Person) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}
