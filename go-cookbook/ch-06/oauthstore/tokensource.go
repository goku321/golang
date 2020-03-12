package oauthstore

import (
	"context"

	"golang.org/x/oauth2"
)

type storageTokenSource struct {
	*Config
	oauth2.TokenSource
}

// Token satisfies the TokenSource interface
func (s *storageTokenSource) Token() (*oauth2.Token, error) {
	if token, err := s.Config.Storage.GetToken(); err == nil && token.Valid() {
		return token, err
	}
	token, err := s.TokenSource.Token()
	if err != nil {
		return token, err
	}
	if err := s.Config.Storage.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

// StorageTokenSource will be user by out configs
// TokenSource function
func StorageTokenSource(ctx context.Context, c *Config, t *oauth2.Token) oauth2.TokenSource {
	return &storageTokenSource{}
}