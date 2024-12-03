package advent

import (
	"bufio"
	"os"
	"strconv"
)

type report struct {
	lines     [][]int
	safeCount int
}

func readFile(path string) *bufio.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(f)
}

func readNumbers(reader *bufio.Reader) ([]int, error) {
	nums := []int{}
	num := []byte{}
	for {
		char, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}

		if char != ' ' {
			num = append(num, char)
		} else {
			nums = append(nums, strconv.Atoi())
		}
	}
}

func parseInput(reader *bufio.Reader) report {

}
