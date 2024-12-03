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

// isSliceSafe returns if a slice of num is safe, which means if it is
// strictly ascending or descending, and if adjacent numbers differ by least MIN_DIFF
// and by most MAX_DIFF
func isSliceSafe(nums []int) bool {
	if len(nums) == 0 {
		panic("Empty nums slice")
	}
	if len(nums) == 1 {
		return true
	}

	left := nums[0]
	right := nums[1]
	ascending := false
	if isInDiffRange(left, right) {
		if left < right {
			ascending = true
		}
	} else {
		return false
	}

	left = nums[1]
	for idx := 2; idx < len(nums); idx++ {
		right := nums[idx]
		if !isInDiffRange(left, right) || !isInOrder(ascending, left, right) {
			return false
		}
		left = right
	}
	return true
}

func parseInput(reader *bufio.Reader) report {
	report := newReport()
	for {
		nums, err := readNumbers(reader)
		if err != nil {
			panic(err)
		}
		isSafe := isSliceSafe(nums)

		report.lines = append(report.lines, nums)
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

	report := parseInput(reader)
	fmt.Printf("The answer is %d\n", report.safeCount)
}
