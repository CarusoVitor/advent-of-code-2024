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

var doRegex = regexp.MustCompile(`do\(\)`)
var dontRegex = regexp.MustCompile(`don't\(\)`)

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
	index int
}

type do struct {
	neg   bool
	index int
}

func newDo(neg bool, index int) do {
	return do{neg, index}
}

func newMult(left, right, index int) mult {
	return mult{left, right, index}
}

func findMultMatches(text string) ([]mult, error) {
	mult := make([]mult, 0, 256)
	matches := multRegex.FindAllStringSubmatch(text, -1)
	indexes := multRegex.FindAllStringSubmatchIndex(text, -1)

	if len(matches) != len(indexes) {
		panic(
			fmt.Sprintf("matches must has len %d different than index (%d)", len(matches), len(indexes)),
		)
	}

	for idx := range matches {
		left := matches[idx][leftIdx]
		leftValue, err := strconv.Atoi(left)
		if err != nil {
			return nil, err
		}

		right := matches[idx][rightIdx]
		rightValue, err := strconv.Atoi(right)
		if err != nil {
			return nil, err
		}
		index := indexes[idx][0]
		multObj := newMult(leftValue, rightValue, index)
		mult = append(mult, multObj)
	}
	return mult, nil
}

func findDoMatches(text string) ([]do, error) {
	do := make([]do, 0, 256)
	indexes := doRegex.FindAllStringSubmatchIndex(text, -1)

	for _, indexSlice := range indexes {
		index := indexSlice[0]
		doObj := newDo(false, index)
		do = append(do, doObj)
	}
	return do, nil
}

func findDontMatches(text string) ([]do, error) {
	dont := make([]do, 0, 256)
	indexes := dontRegex.FindAllStringSubmatchIndex(text, -1)

	for _, indexSlice := range indexes {
		index := indexSlice[0]
		dontObj := newDo(true, index)
		dont = append(dont, dontObj)
	}
	return dont, nil
}

func parseInput(text string) ([]mult, []do, []do) {
	mults, err := findMultMatches(text)
	if err != nil {
		panic(err)
	}
	do, err := findDoMatches(text)
	if err != nil {
		panic(err)
	}
	dont, err := findDontMatches(text)
	if err != nil {
		panic(err)
	}

	return mults, do, dont
}

func sumMultsResults(mults []mult) int {
	sum := 0
	for _, multObj := range mults {
		sum += multObj.left * multObj.right
	}
	return sum
}

func MultiOverPartOne() {
	path := "advent/3multitover/input.txt"
	text := readFile(path)

	mults, _, _ := parseInput(text)
	result := sumMultsResults(mults)
	fmt.Printf("The answer is %d\n", result)
}
