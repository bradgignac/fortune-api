package fortune

import (
	"crypto/sha1"
	"fmt"
	"io"
)

// Fortune is a data structure representing a fortune.
type Fortune struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

// NewFortune instantiates a fortune.Fortune from the provided string.
func NewFortune(data string) Fortune {
	return Fortune{ID: ComputeID(data), Data: data}
}

func ComputeID(data string) string {
	sha := sha1.New()
	io.WriteString(sha, data)
	bytes := sha.Sum(nil)
	id := fmt.Sprintf("%x", bytes)
	return id[0:7]
}
