package main

import (
	"encoding/json"
	"time"

	"log"
)

// Account holds account details.
type Account struct {
	ID            string `json:"id"`
	Object        string `json:"object"`
	RoutingNumber string `json:"routing_number"`
}

// Card holds card details
type Card struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Last4  string `json:"last4"`
}

// Data holds card and account details.
type Data struct {
	*Card
	*Account
}

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
func (t Time) UnmarshalJSON(data []byte) error {
	var i int64
	if err := json.Unmarshal(data, &i); err != nil {
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

	jsonString = `
	{
	  "data": {
		"object": "card",
		"id": "card_123",
		"last4": "4242"
	  }
	}
	`

	px := Person{}
	var m map[string]interface{}
	if b, err := json.Marshal(px); err != nil {
		log.Fatal(err)
	} else {
		if err := json.Unmarshal(b, &m); err != nil {
			log.Fatal(err)
		}
	}

	jsonString = `
		{
			"data": {
			"object": "card",
			"id": "card_123",
			"last4": "4242"
			}
	  	}
	`

	var dm map[string]Data
	if err := json.Unmarshal([]byte(jsonString), &dm); err != nil {
		log.Fatal(err)
	}

	data := dm["data"]
	if data.Card != nil {
		log.Println(data.Card)
	}
	if data.Account != nil {
		log.Println(data.Account)
	}

	b, err = json.Marshal(dm)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}

// MarshalJSON marshals Data.
func (d Data) MarshalJSON() ([]byte, error) {
	if d.Account != nil {
		return json.Marshal(d.Account)
	} else if d.Card != nil {
		return json.Marshal(d.Card)
	} else {
		return json.Marshal(nil)
	}
}

// UnmarshalJSON un-marshals json into Data.
func (d Data) UnmarshalJSON(data []byte) error {
	temp := struct {
		Object string `json:"object"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
}
