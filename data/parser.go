package data

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// Parse initializes a fortune.Database with the fortune data parsed from the provided io.Reader.
func Parse(r io.Reader) (*Database, error) {
	fortunes := make([]string, 0)
	scanner := bufio.NewScanner(r)
	scanner.Split(splitFortune)

	for scanner.Scan() {
		fortunes = append(fortunes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return NewDatabase(fortunes), nil
}

func splitFortune(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	const escape = "\n%\n"
	const advance = len(escape)
	if i := bytes.Index(data, []byte(escape)); i >= 0 {
		return i + advance, trimNewLines(data[0:i]), nil
	}

	if atEOF {
		return len(data), trimNewLines(data), nil
	}

	return 0, nil, nil
}

func trimNewLines(data []byte) []byte {
	trimmed := strings.Trim(string(data), "\n")
	return []byte(trimmed)
}
