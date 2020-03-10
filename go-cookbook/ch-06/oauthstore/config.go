package oauthstore

import (
	"context"

	"golang.org/x/oauth2"
)

// Config wraps the default oauth2.Config
// and adds storage
type Config struct {
	*oauth2.Config
	Storage
}

// Exchange stores a token after retrieval
func (c *Config) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	if err := c.Storage.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}
