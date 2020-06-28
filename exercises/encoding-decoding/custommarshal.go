package main

import (
	"encoding/json"
	"time"

	"log"
)

// Person represents a person.
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
	BornAt    time.Time
}

// JSONPerson represents Marshalled Person.
type JSONPerson struct {
	FirstName, LastName string
	Age                 int
	Gender              string
	BornAt              int64
}

func newJSONPerson(p Person) JSONPerson {
	return JSONPerson{
		p.FirstName,
		p.LastName,
		p.Age,
		p.Gender,
		p.BornAt.Unix(),
	}
}

// MarshalJSON encodes Person to JSON.
func (p *Person) MarshalJSON() ([]byte, error) {
	jsonP := newJSONPerson(*p)
	return json.Marshal(&jsonP)
}

func main() {
	p := &Person{"Go", "Lang", 11, "L", time.Date(2009, 10, 1, 0, 0, 0, 0, time.UTC)}
	b, err := json.Marshal(&p)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}
