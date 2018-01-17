package fortune

import (
	"math/rand"
)

// Database is an in-memory container for fortunes.
type Database struct {
	data map[string]Fortune
}

// NewDatabase creates a fortune.Database from an array of fortunes.
func NewDatabase(fortunes []string) *Database {
	data := make(map[string]Fortune, len(fortunes))
	for _, raw := range fortunes {
		parsed := NewFortune(raw)
		data[parsed.ID] = parsed
	}
	return &Database{data: data}
}

// List returns all fortunes in the database.
func (d *Database) List() []Fortune {
	vals := make([]Fortune, 0)
	for _, v := range d.data {
		vals = append(vals, v)
	}
	return vals
}

// Get returns a single fortune from the database.
func (d *Database) Get(id string) Fortune {
	return d.data[id]
}

// Random returns the ID of a random fortune in the database.
func (d *Database) Random() string {
	idx := rand.Intn(len(d.data))
	keys := make([]string, 0)
	for k := range d.data {
		keys = append(keys, k)
	}
	return keys[idx]
}

// Count returns the number of fortunes in the database.
func (d *Database) Count() int {
	return len(d.data)
}
