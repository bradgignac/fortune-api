package db

import (
	"math/rand"

	"github.com/juju/errors"
)

// ErrEmptyDatabase is returned when an empty database is used.
var ErrEmptyDatabase = errors.New("Database is empty")

// ErrMissingFortune is returned when a fortune for the provided ID is not found.
var ErrMissingFortune = errors.New("Could not find fortune with provided ID")

// Database is an in-memory container for fortunes.
type Database struct {
	data map[string]*Fortune
}

// NewDatabase creates a fortune.Database from an array of fortunes.
func NewDatabase(fortunes []string) *Database {
	data := make(map[string]*Fortune, len(fortunes))
	for _, raw := range fortunes {
		parsed := NewFortune(raw)
		data[parsed.ID] = &parsed
	}
	return &Database{data: data}
}

// List returns all fortunes in the database.
func (d *Database) List() []*Fortune {
	vals := make([]*Fortune, 0)
	for _, v := range d.data {
		vals = append(vals, v)
	}
	return vals
}

// Get returns a single fortune from the database.
func (d *Database) Get(id string) (*Fortune, error) {
	if f, ok := d.data[id]; ok {
		return f, nil
	}

	return nil, ErrMissingFortune
}

// Random returns the ID of a random fortune in the database.
func (d *Database) Random() (string, error) {
	if len(d.data) == 0 {
		return "", ErrEmptyDatabase
	}

	idx := rand.Intn(len(d.data))
	keys := make([]string, 0)
	for k := range d.data {
		keys = append(keys, k)
	}
	return keys[idx], nil
}

// Count returns the number of fortunes in the database.
func (d *Database) Count() int {
	return len(d.data)
}
