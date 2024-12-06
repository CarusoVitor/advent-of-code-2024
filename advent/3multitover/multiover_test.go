package advent

import (
	"reflect"
	"testing"
)

func TestFindMultMatches(t *testing.T) {
	text := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := []mult{
		newMult(2, 4, 1),
		newMult(5, 5, 29),
		newMult(11, 8, 53),
		newMult(8, 5, 62),
	}
	result, err := findMultMatches(text)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("result %#v != expected %#v", result, expected)
	}
}

func TestFindDoMatches(t *testing.T) {
	text := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := []do{
		newDo(false, 59),
	}
	result, err := findDoMatches(text)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("result %#v != expected %#v", result, expected)
	}
}

func TestFindDontMatches(t *testing.T) {
	text := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := []do{
		newDo(true, 20),
	}
	result, err := findDontMatches(text)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("result %#v != expected %#v", result, expected)
	}
}
