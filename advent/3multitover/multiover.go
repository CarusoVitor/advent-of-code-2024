package advent

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var multRegex = regexp.MustCompile(`mul\((?P<left>\d+),(?P<right>\d+)\)`)
var leftIdx = multRegex.SubexpIndex("left")
var rightIdx = multRegex.SubexpIndex("right")

func readFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	text, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(text)
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

func findMultMatches(text string) ([]mult, error) {
	mult := make([]mult, 0, 256)
	matches := multRegex.FindAllStringSubmatch(text, -1)

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

func parseInput(text string) []mult {
	matches, err := findMultMatches(text)
	if err != nil {
		panic(err)
	}
	return matches
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
	text := readFile(path)

	mults := parseInput(text)
	result := sumMultsResults(mults)
	fmt.Printf("The answer is %d\n", result)
}
