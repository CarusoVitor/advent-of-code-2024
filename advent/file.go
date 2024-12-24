package advent

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(path string) *bufio.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(f)
}

func NewTestReader(text string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(text))
}
