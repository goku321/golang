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
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int
	Gender    string
	BornAt    int64
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

// UnmarshalJSON decodes JSON object to Person.
func (p *Person) UnmarshalJSON(data []byte) error {
	var jsonP *JSONPerson
	err := json.Unmarshal(data, &jsonP)
	if err != nil {
		return err
	}
	p.BornAt = time.Unix(jsonP.BornAt, 0)
	p.Age = jsonP.Age
	p.FirstName = jsonP.FirstName
	p.LastName = jsonP.LastName
	p.Gender = jsonP.Gender

	return nil
}

func main() {
	p := &Person{"Go", "Lang", 11, "L", time.Date(2009, 10, 1, 0, 0, 0, 0, time.UTC)}
	b, err := json.Marshal(&p)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))

	jsonString := `
	{
		"first_name": "Go",
		"last_name": "Lang",
		"gender": "L",
		"age": 11,
		"born_at": 1254355200
	}
	`

	var newPerson *Person
	err = json.Unmarshal([]byte(jsonString), &newPerson)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newPerson)
}
