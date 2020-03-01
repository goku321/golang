package user

import "fmt"

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func (u *User) Info() string {
	return fmt.Sprintf("%s %s is %d years old", u.FirstName, u.LastName, u.Age)
}
