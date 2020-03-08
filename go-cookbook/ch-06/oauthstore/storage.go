package oauthstore

import (
	"golang.org/x/oauth2"
)

// Storage is a generic storage interface
type Storage interface {
	GetToken() (*oauth2.Token, error)
	SetToken(*oauth2.Token) error
}
