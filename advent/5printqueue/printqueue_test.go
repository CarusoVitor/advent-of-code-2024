package advent

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func newReader(text string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(text))
}

func TestReadRules(t *testing.T) {
	rd := newReader("47|5\n" +
		"97|13\n" +
		"97|61\n" +
		"97|47\n\n")
	expected := rule{
		5:  {47},
		13: {97},
		61: {97},
		47: {97},
	}
	result := readRules(rd)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %#v != result %#v", expected, result)
	}
}

func TestReadPages(t *testing.T) {
	rd := newReader("1,2,3,4\n5,6,7,8,9,10")
	expected := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8, 9, 10},
	}
	result := readPages(rd)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %#v != result %#v", expected, result)
	}
}

func TestValidPage(t *testing.T) {
	rules := rule{
		5:  {47},
		13: {97},
		61: {97},
		47: {97},
	}
	page := []int{97, 47, 61}
	valid := isValidPage(page, rules)
	if !valid {
		t.Errorf("page %#v must be valid", page)
	}
}

func TestInvalidPage(t *testing.T) {
	rules := rule{
		5:  {47},
		13: {97},
		61: {97},
		47: {97},
	}
	page := []int{13, 97, 47, 61}
	valid := isValidPage(page, rules)
	if valid {
		t.Errorf("page %#v must be invalid", page)
	}
}

func TestPrintQueuePartOne(t *testing.T) {
	input := "47|53\n" +
		"97|13\n" +
		"97|61\n" +
		"97|47\n" +
		"75|29\n" +
		"61|13\n" +
		"75|53\n" +
		"29|13\n" +
		"97|29\n" +
		"53|29\n" +
		"61|53\n" +
		"97|53\n" +
		"61|29\n" +
		"47|13\n" +
		"75|47\n" +
		"97|75\n" +
		"47|61\n" +
		"75|61\n" +
		"47|29\n" +
		"75|13\n" +
		"53|13\n" +
		"\n" +
		"75,47,61,53,29\n" +
		"97,61,53,29,13\n" +
		"75,29,13\n" +
		"75,97,47,61,53\n" +
		"61,13,29\n" +
		"97,13,75,29,47"
	rd := newReader(input)

	expected := 143
	result := PrintQueuePartOne(rd)

	if result != expected {
		t.Errorf("expected %d != result %d", expected, result)
	}
}

func TestPrintQueuePartTwo(t *testing.T) {
	input := "47|53\n" +
		"97|13\n" +
		"97|61\n" +
		"97|47\n" +
		"75|29\n" +
		"61|13\n" +
		"75|53\n" +
		"29|13\n" +
		"97|29\n" +
		"53|29\n" +
		"61|53\n" +
		"97|53\n" +
		"61|29\n" +
		"47|13\n" +
		"75|47\n" +
		"97|75\n" +
		"47|61\n" +
		"75|61\n" +
		"47|29\n" +
		"75|13\n" +
		"53|13\n" +
		"\n" +
		"75,47,61,53,29\n" +
		"97,61,53,29,13\n" +
		"75,29,13\n" +
		"75,97,47,61,53\n" +
		"61,13,29\n" +
		"97,13,75,29,47"
	rd := newReader(input)

	expected := 123
	result := PrintQueuePartTwo(rd)

	if result != expected {
		t.Errorf("expected %d != result %d", expected, result)
	}
}
