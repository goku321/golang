package user

import "fmt"

type Client struct {
}

var printName func(name string)

func NewClient() *Client {
	printName = func(name string) {
		fmt.Println(name)
	}
	return &Client{}
}
