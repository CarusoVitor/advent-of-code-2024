package advent

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type rule map[int][]int

func (r rule) add(left, right int) {
	if _, ok := r[right]; !ok {
		r[right] = make([]int, 0, 32)
	}
	r[right] = append(r[right], left)
}

// readRules read all lines in the format XX|YY
// storing in a map that maps the succeding numbers
// to all its predecessors
// For example:
// 15|22     22: {15,31}
// 31/22 ==> 31: {15}
// 15/31
func readRules(rd *bufio.Reader) rule {
	rules := make(rule)
	for {
		line, err := rd.ReadString('\n')
		line = line[:len(line)-1]
		if err != nil {
			panic(err)
		}
		// pages end is a single "\n"
		if len(line) == 0 {
			break
		}
		nums := strings.Split(line, "|")
		if len(nums) != 2 {
			panic(fmt.Sprintf("rule should have 2 numbers: %s", line))
		}
		left, err := strconv.Atoi(nums[0])
		right, err := strconv.Atoi(nums[1])
		rules.add(left, right)
	}
	return rules
}

func stringSliceToInt(slice []string) ([]int, error) {
	nums := make([]int, len(slice))
	for idx, entry := range slice {
		num, err := strconv.Atoi(entry)
		if err != nil {
			return nil, err
		}
		nums[idx] = num
	}
	return nums, nil
}

func readPages(rd *bufio.Reader) [][]int {
	pages := make([][]int, 0, 256)
	for {
		line, fileErr := rd.ReadString('\n')
		if fileErr != io.EOF {
			if fileErr != nil {
				panic(fileErr)
			}
			line = line[:len(line)-1]
		}
		nums := strings.Split(line, ",")
		page, err := stringSliceToInt(nums)
		if err != nil {
			panic(err)
		}
		pages = append(pages, page)
		if fileErr == io.EOF {
			break
		}
	}
	return pages
}

func parse(rd *bufio.Reader) (rule, [][]int) {
	rules := readRules(rd)
	pages := readPages(rd)
	return rules, pages
}

func isValidPage(page []int, rules rule) bool {
	mustNotAppear := make(map[int]bool)
	for _, num := range page {
		if mustNotAppear[num] {
			return false
		}
		for _, invalid := range rules[num] {
			mustNotAppear[invalid] = true
		}
	}
	return true
}

func getValidPages(pages [][]int, rules rule) [][]int {
	validPages := make([][]int, 0, len(pages))
	for _, page := range pages {
		if isValidPage(page, rules) {
			validPages = append(validPages, page)
		}
	}
	return validPages
}

func getInvalidPages(pages [][]int, rules rule) [][]int {
	invalidPages := make([][]int, 0, len(pages))
	for _, page := range pages {
		if !isValidPage(page, rules) {
			invalidPages = append(invalidPages, page)
		}
	}
	return invalidPages
}
func sumMiddleNumber(nums [][]int) int {
	sum := 0
	for _, slice := range nums {
		mid := (len(slice) - 1) / 2
		sum += slice[mid]
	}
	return sum
}

func PrintQueuePartOne(rd *bufio.Reader) int {
	rules, pages := parse(rd)
	validPages := getValidPages(pages, rules)
	sum := sumMiddleNumber(validPages)
	return sum
}

func in(num int, slice []int) bool {
	for _, n := range slice {
		if num == n {
			return true
		}
	}
	return false
}

func reorderInvalidPages(invalidPages [][]int, rules rule) [][]int {
	for _, page := range invalidPages {
		slices.SortFunc(page, func(a, b int) int {
			if a == b {
				return 0
			}
			beforeB := rules[b]
			if in(a, beforeB) {
				return 1
			}
			beforeA := rules[a]
			if in(b, beforeA) {
				return -1
			}
			return 0
		})
	}
	return invalidPages
}

func PrintQueuePartTwo(rd *bufio.Reader) int {
	rules, pages := parse(rd)
	invalidPages := getInvalidPages(pages, rules)
	reordered := reorderInvalidPages(invalidPages, rules)
	sum := sumMiddleNumber(reordered)
	return sum
}
