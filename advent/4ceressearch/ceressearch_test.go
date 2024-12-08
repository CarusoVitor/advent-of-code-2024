package advent

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestParseInputXPositions(t *testing.T) {
	input := "MMMSXXMASM\n" +
		"MSAMXMSMSA\n"
	rd := bufio.NewReader(strings.NewReader(input))
	cross := newCrossWordFromReader(rd)

	expectedXPositions := [][]int{{4, 5}, {4}}

	if !reflect.DeepEqual(cross.xPositions, expectedXPositions) {
		t.Errorf("expected %#v != result #%v", expectedXPositions, cross.xPositions)
	}
}

func TestParseInputWidthHeigth(t *testing.T) {
	input := "MMMS\n" +
		"MSAM\n" +
		"MSAM\n"
	rd := bufio.NewReader(strings.NewReader(input))
	cross := newCrossWordFromReader(rd)

	expectedWidth := 4
	expectedHeigth := 3

	if cross.width != expectedWidth {
		t.Errorf("expected width %d != result width %d", expectedWidth, cross.width)
	}

	if cross.heigth != expectedHeigth {
		t.Errorf("expected heigth %d != result heigth %d", expectedHeigth, cross.heigth)
	}
}

func TestIsXMASPresent(t *testing.T) {
	input := "XMAS\n"
	rd := bufio.NewReader(strings.NewReader(input))
	cross := newCrossWordFromReader(rd)
	expected := 1
	result := cross.allXMASCount()
	if result != expected {
		t.Errorf("expected %d != result %d", expected, result)
	}
}

func TestIsXMASPresentMultipleXSameLine(t *testing.T) {
	input :=
		"XXMAS\n" +
			"MM...\n" +
			"AA...\n" +
			"SS...\n"
	rd := bufio.NewReader(strings.NewReader(input))
	cross := newCrossWordFromReader(rd)
	expected := 3
	result := cross.allXMASCount()
	if result != expected {
		t.Errorf("expected %d != result %d", expected, result)
	}
}

func TestIsOneXMASPresentMultipleAllPossibilites(t *testing.T) {
	input :=
		"S..S..S\n" +
			".A.A.A.\n" +
			"..MMM..\n" +
			"SAMXMAS\n" +
			"..MMM..\n" +
			".A.A.A.\n" +
			"S..S..S\n"
	rd := bufio.NewReader(strings.NewReader(input))
	cross := newCrossWordFromReader(rd)
	expected := 8
	result := cross.allXMASCount()
	if result != expected {
		t.Errorf("expected %d != result %d", expected, result)
	}
}

func TestIsMultipleXMASPresentMultipleAllPossibilites(t *testing.T) {
	input :=
		"S..S..S...\n" +
			".A.A.A....\n" +
			"..MMM.....\n" +
			"SAMXMAS..S\n" +
			"..MMM....A\n" +
			".A.A.A...M\n" +
			"S..S..SAMX\n"
	rd := bufio.NewReader(strings.NewReader(input))
	cross := newCrossWordFromReader(rd)
	expected := 10
	result := cross.allXMASCount()
	if result != expected {
		t.Errorf("expected %d != result %d", expected, result)
	}
}

func TestXmasCountSampleInput(t *testing.T) {
	input :=
		"....XXMAS.\n" +
			".SAMXMS...\n" +
			"...S..A...\n" +
			"..A.A.MS.X\n" +
			"XMASAMX.MM\n" +
			"X.....XA.A\n" +
			"S.S.S.S.SS\n" +
			".A.A.A.A.A\n" +
			"..M.M.M.MM\n" +
			".X.X.XMASX\n"
	rd := bufio.NewReader(strings.NewReader(input))
	cross := newCrossWordFromReader(rd)
	expected := 18
	result := cross.allXMASCount()
	if result != expected {
		t.Errorf("expected %d != result %d", expected, result)
	}

}

func TestMultipleXmasToRight(t *testing.T) {
	input := "XMASMSXSSSXMASMXXSAMXSMMSSSMXSXSXSMMSXMAS\n"
	rd := bufio.NewReader(strings.NewReader(input))
	cross := newCrossWordFromReader(rd)
	expected := 3
	result := cross.allXMASCount()
	if result != expected {
		t.Errorf("expected %d != result %d", expected, result)
	}
}
