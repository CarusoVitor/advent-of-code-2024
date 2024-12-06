package advent

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestFindMultMatches(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"))
	expected := []mult{
		newMult(2, 4),
		newMult(5, 5),
		newMult(11, 8),
		newMult(8, 5),
	}
	result := parseInput(scanner)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("result %#v != expected %#v", result, expected)
	}
}
