package advent

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var multRegex = regexp.MustCompile(`mul\((?P<left>\d+),(?P<right>\d+)\)`)
var leftIdx = multRegex.SubexpIndex("left")
var rightIdx = multRegex.SubexpIndex("right")

func readFile(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}

type mult struct {
	left  int
	right int
	value int
}

func newMult(left, right int) mult {
	value := left * right
	return mult{left, right, value}
}

func findMultMatches(line string) ([]mult, error) {
	mult := make([]mult, 0, 256)
	matches := multRegex.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		left := match[leftIdx]
		leftValue, err := strconv.Atoi(left)
		if err != nil {
			return nil, err
		}

		right := match[rightIdx]
		rightValue, err := strconv.Atoi(right)
		if err != nil {
			return nil, err
		}
		multObj := newMult(leftValue, rightValue)
		mult = append(mult, multObj)
	}
	return mult, nil
}

func parseInput(scanner *bufio.Scanner) []mult {
	ok := scanner.Scan()
	mult := make([]mult, 0, 256)
	for ok {
		line := scanner.Text()

		matches, err := findMultMatches(line)
		if err != nil {
			panic(err)
		}

		mult = append(mult, matches...)
		ok = scanner.Scan()
	}
	return mult
}

func sumMultsResults(mults []mult) int {
	sum := 0
	for _, multObj := range mults {
		sum += multObj.value
	}
	return sum
}

func MultiOverPartOne() {
	path := "advent/3multitover/input.txt"
	scanner := readFile(path)

	mults := parseInput(scanner)
	result := sumMultsResults(mults)
	fmt.Printf("The answer is %d\n", result)
}
