package oauthstore

import (
	"encoding/json"
	"os"
	"sync"

	"golang.org/x/oauth2"
)

// FileStorage satisfies our storage interface
type FileStorage struct {
	Path string
	mu   sync.RWMutex
}

// GetToken retrieves a token from a file
func (f *FileStorage) GetToken() (*oauth2.Token, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	in, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	var t *oauth2.Token
	data := json.NewDecoder(in)
	return t, data.Decode(&t)
}
