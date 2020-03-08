package oauthstore

import "golang.org/x/oauth2"

// Config wraps the default oauth2.Config
// and adds storage
type Config struct {
	*oauth2.Config
	Storage
}
