package advent

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const LIST_SIZE = 1028

type locationIDs struct {
	left       []int
	right      []int
	rightCount map[int]int
}

func newLocationId() locationIDs {
	left := make([]int, 0, LIST_SIZE)
	right := make([]int, 0, LIST_SIZE)
	rightCount := make(map[int]int, LIST_SIZE)

	return locationIDs{left, right, rightCount}
}

func (loc *locationIDs) diffSortedSlices() int {
	leftSorted := make([]int, len(loc.left))
	copy(leftSorted, loc.left)

	rightSorted := make([]int, len(loc.right))
	copy(rightSorted, loc.right)

	sort.Slice(leftSorted, func(i, j int) bool {
		return leftSorted[i] > leftSorted[j]
	})

	sort.Slice(rightSorted, func(i, j int) bool {
		return rightSorted[i] > rightSorted[j]
	})

	sum := 0
	for idx := range leftSorted {
		sum += absDiff(leftSorted[idx], rightSorted[idx])
	}
	return sum
}

func (loc *locationIDs) multiplyLeftOnRightCount() int {
	sum := 0
	for _, num := range loc.left {
		sum += num * loc.rightCount[num]
	}
	return sum
}

func (loc *locationIDs) addLeft(num int) {
	loc.left = append(loc.left, num)
}

func (loc *locationIDs) addRight(num int) {
	loc.right = append(loc.right, num)
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func readFile(path string) *bufio.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(f)
}

func getLineNums(reader *bufio.Reader) (int, int, error) {
	line, _, err := reader.ReadLine()
	if err != nil {
		return 0, 0, err
	}
	numbers := strings.Split(string(line), "   ")
	if len(numbers) != 2 {
		return 0, 0, fmt.Errorf("line must have exactly 2 numbers, got %d", len(numbers))
	}

	leftNum, err := strconv.Atoi(numbers[0])
	if err != nil {
		return 0, 0, err
	}
	rightNum, err := strconv.Atoi(numbers[1])
	if err != nil {
		return 0, 0, err
	}
	return leftNum, rightNum, nil
}

func parseInput(reader *bufio.Reader) locationIDs {
	locations := newLocationId()

	for {
		leftNum, rightNum, err := getLineNums(reader)
		if err != nil {
			panic(err)
		}

		locations.addRight(rightNum)
		locations.addLeft(leftNum)

		locations.rightCount[rightNum]++

		_, err = reader.Peek(1)
		if err == io.EOF {
			break
		}
	}

	return locations
}

func HistorianHysteriaPartOne() {
	path := "advent/1historianhysteria/input.txt"
	reader := readFile(path)

	locations := parseInput(reader)
	sum := locations.diffSortedSlices()
	fmt.Printf("[1] The answer is %d\n", sum)
}

func HistorianHysteriaPartTwo() {
	path := "advent/1historianhysteria/input.txt"
	reader := readFile(path)

	locations := parseInput(reader)
	sum := locations.multiplyLeftOnRightCount()
	fmt.Printf("[2] The answer is %d\n", sum)
}
