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

func findRegexIndexes(text string, regex *regexp.Regexp) ([]int, error) {
	do := make([]int, 0, 256)
	indexes := regex.FindAllStringSubmatchIndex(text, -1)

	for _, indexSlice := range indexes {
		index := indexSlice[0]
		do = append(do, index)
	}
	return do, nil
}

func parseInput(text string) ([]mult, []int, []int) {
	mults, err := findMultMatches(text)
	if err != nil {
		panic(err)
	}
	do, err := findRegexIndexes(text, doRegex)
	if err != nil {
		panic(err)
	}
	dont, err := findRegexIndexes(text, dontRegex)
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

// incrementUntilCloseTo increments the index until [nums[index]] is as close
// as possible to goal, but less or equal than goal. It assumes nums is an ascending slice
func incrementUntilCloseTo(nums []int, idx, goal int) int {
	if idx == len(nums)-1 {
		return idx
	}
	for idx < len(nums) {
		if nums[idx] > goal {
			return idx - 1
		}
		idx++
	}
	return idx - 1
}

func sumValidMults(mults []mult, doIndexes []int, dontIndexes []int) int {
	sum := 0
	doIdx := 0
	dontIdx := 0
	for _, mult := range mults {
		if doIndexes[doIdx] < mult.index {
			doIdx = incrementUntilCloseTo(doIndexes, doIdx, mult.index)
		}
		if dontIndexes[dontIdx] < mult.index {
			dontIdx = incrementUntilCloseTo(dontIndexes, dontIdx, mult.index)
		}

		isLastDo := doIndexes[doIdx] > dontIndexes[dontIdx] && doIndexes[doIdx] < mult.index
		noMultsToLeft := dontIndexes[dontIdx] > mult.index

		if isLastDo || noMultsToLeft {
			sum += mult.left * mult.right
		}
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

func MultiOverPartTwo() {
	path := "advent/3multitover/input.txt"
	text := readFile(path)

	mults, doIndexes, dontIndexes := parseInput(text)
	result := sumValidMults(mults, doIndexes, dontIndexes)
	fmt.Printf("The answer is %d\n", result)
}
