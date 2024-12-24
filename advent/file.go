package advent

import (
	"bufio"
	"os"
)

func ReadFile(path string) *bufio.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(f)
}
