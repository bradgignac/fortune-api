package fortune

import (
	"math/rand"
)

type Database struct {
	data map[string]Fortune
}

func NewDatabase(fortunes []string) *Database {
	data := make(map[string]Fortune, len(fortunes))
	for _, raw := range fortunes {
		parsed := NewFortune(raw)
		data[parsed.Id] = parsed
	}
	return &Database{data: data}
}

func (d *Database) List() []Fortune {
	vals := make([]Fortune, 0)
	for _, v := range d.data {
		vals = append(vals, v)
	}
	return vals
}

func (d *Database) Get(id string) Fortune {
	return d.data[id]
}

func (d *Database) Random() string {
	idx := rand.Intn(len(d.data))
	keys := make([]string, 0)
	for k := range d.data {
		keys = append(keys, k)
	}
	return keys[idx]
}

func (d *Database) Count() int {
	return len(d.data)
}
