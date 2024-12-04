package advent

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const MAX_DIFF = 3
const MIN_DIFF = 1

type report struct {
	lines     [][]int
	safeCount int
}

func newReport() report {
	lines := make([][]int, 0, 1028)
	return report{lines: lines}
}

func readFile(path string) *bufio.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(f)
}

func readNumbers(reader *bufio.Reader) ([]int, error) {
	line, _, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}
	numsStr := strings.Fields(string(line))
	nums := make([]int, len(numsStr))

	for idx, num := range numsStr {
		numValue, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		nums[idx] = numValue
	}

	return nums, nil
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func isInDiffRange(num1, num2 int) bool {
	diff := absDiff(num1, num2)
	return diff >= MIN_DIFF && diff <= MAX_DIFF
}

func isInOrder(ascending bool, num1, num2 int) bool {
	if ascending {
		return num1 < num2
	}
	return num1 > num2
}

// checkSliceSafe returns if a slice of num is safe, which means if it is
// strictly ascending or descending, and if adjacent numbers differ by least MIN_DIFF
// and by most MAX_DIFF. In unsafe sequencies it returns the indexes of the first numbers
// pair that make the sequence unsafe
func checkSliceSafe(nums []int) (bool, int, int) {
	if len(nums) == 0 {
		panic("Empty nums slice")
	}
	if len(nums) == 1 {
		return true, -1, -1
	}

	left := nums[0]
	right := nums[1]
	ascending := false
	if isInDiffRange(left, right) {
		if left < right {
			ascending = true
		}
	} else {
		return false, 0, 1
	}

	left = nums[1]
	for idx := 2; idx < len(nums); idx++ {
		right := nums[idx]
		if !isInDiffRange(left, right) || !isInOrder(ascending, left, right) {
			return false, idx - 1, idx
		}
		left = right
	}
	return true, -1, -1
}

func removeIndex(s []int, idx int) []int {
	ret := make([]int, 0, len(s)-1)
	ret = append(ret, s[:idx]...)
	return append(ret, s[idx+1:]...)
}

// isSliceSafe checks if a slice is safe, considering one or two chances
// in case of two chances, the only possible indexes to be removed are the two that
// the comparison went wrong (left and right index), and the 0th index, since it's
// the one that may be breaking the sequence order
func isSliceSafe(nums []int, oneMoreChance bool) bool {
	isSafe, leftIdx, rightIdx := checkSliceSafe(nums)
	if oneMoreChance && !isSafe {
		isSafeWithoutLeft, _, _ := checkSliceSafe(removeIndex(nums, leftIdx))
		isSafeWithoutRight, _, _ := checkSliceSafe(removeIndex(nums, rightIdx))
		isSafeWithoutStart, _, _ := checkSliceSafe(removeIndex(nums, 0))
		isSafe = isSafeWithoutLeft || isSafeWithoutRight || isSafeWithoutStart
	}
	return isSafe
}

func parseInput(reader *bufio.Reader, oneMoreChance bool) report {
	report := newReport()
	for {
		nums, err := readNumbers(reader)
		if err != nil {
			panic(err)
		}

		report.lines = append(report.lines, nums)
		isSafe := isSliceSafe(nums, oneMoreChance)
		if isSafe {
			report.safeCount++
		}

		_, err = reader.Peek(1)
		if err == io.EOF {
			break
		}
	}
	return report
}

func RedNosedPartOne() {
	path := "advent/2rednosed/input.txt"
	reader := readFile(path)

	report := parseInput(reader, false)
	fmt.Printf("The answer is %d\n", report.safeCount)
}

func RedNosedPartTwo() {
	path := "advent/2rednosed/input.txt"
	reader := readFile(path)

	report := parseInput(reader, true)
	fmt.Printf("The answer is %d\n", report.safeCount)
}
