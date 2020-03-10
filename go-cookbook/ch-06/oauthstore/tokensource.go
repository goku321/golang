package oauthstore

import (
	"context"

	"golang.org/x/oauth2"
)

type storageTokenSource struct {
	*Config
	oauth2.TokenSource
}

// StorageTokenSource will be user by out configs
// TokenSource function
func StorageTokenSource(ctx context.Context, c *Config, t *oauth2.Token) oauth2.TokenSource {
	return &storageTokenSource{}
}