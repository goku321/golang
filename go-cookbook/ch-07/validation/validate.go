package validation

import "errors"

// Verror ...
type Verror struct {
	error
}

// Payload ...
type Payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ValidatePayload ...
func ValidatePayload(p *Payload) error {
	if p.Name == "" {
		return Verror{errors.New("name is required")}
	}

	if p.Age <= 0 && p.Age >= 120 {
		return Verror{errors.New("age is required and must be greater than 0 and less than 120")}
	}
	return nil
}
