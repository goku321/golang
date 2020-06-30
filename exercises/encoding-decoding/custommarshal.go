package main

import (
	"encoding/json"
	"time"

	"log"
)

// Person represents a person.
type Person struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	BornAt    time.Time `json:"born_at"`
}

// PersonX is an alias to Person.
type PersonX Person

// Time embeds time.Time to enable custom (un)marshalling.
type Time struct {
	time.Time
}

// MarshalJSON implements custom marshaler for Time.
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Unix())
}

// UnmarshalJSON implements custom unmarshaler for Time.
func (t Time) UnmarshalJSON([]byte) error {
	var i int64
	if err := json.Unmarshal(&i); err != nil {
		return err
	}
	t.Time = time.Unix(i, 0)
	return nil
}

// JSONPerson represents Marshalled Person.
type JSONPerson struct {
	PersonX
	BornAt int64 `json:"born_at"`
}

func newJSONPerson(p Person) JSONPerson {
	return JSONPerson{
		PersonX(p),
		p.BornAt.Unix(),
	}
}

// ToPerson converts JSONPerson to Person.
func (jp JSONPerson) ToPerson() Person {
	p := Person(jp.PersonX)
	p.BornAt = time.Unix(jp.BornAt, 0)
	return p
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

	*p = jsonP.ToPerson()
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
