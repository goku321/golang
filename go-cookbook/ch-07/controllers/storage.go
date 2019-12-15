package controllers

// Storage interface supports Get and Put
type Storage interface {
	Get() string
	Put(string)
}

// MemStorage implements storage
type MemStorage struct {
	value string
}

// Get in-memory value
func (m *MemStorage) Get() string {
	return m.value
}

// Put in-memory value
func (m *MemStorage) Put(s string) {
	m.value = s
}