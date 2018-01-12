package fortune

import (
	"crypto/sha1"
	"fmt"
	"io"
)

type Fortune struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

func NewFortune(data string) Fortune {
	return Fortune{Id: computeId(data), Data: data}
}

func computeId(data string) string {
	sha := sha1.New()
	io.WriteString(sha, data)
	bytes := sha.Sum(nil)
	id := fmt.Sprintf("%x", bytes)
	return id[0:7]
}
