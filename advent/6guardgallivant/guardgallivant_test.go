package advent

import (
	"testing"

	"github.com/CarusoVitor/advent-of-code-2024/advent"
)

func TestGuardGallivantPartOne(t *testing.T) {
	input := "....#.....\n" +
		".........#\n" +
		"..........\n" +
		"..#.......\n" +
		".......#..\n" +
		"..........\n" +
		".#..^.....\n" +
		"........#.\n" +
		"#.........\n" +
		"......#..."
	rd := advent.NewTestReader(input)

	expected := 41
	result := GuardGallivantPartOne(rd)

	if expected != result {
		t.Errorf("expected %d != result %d", expected, result)
	}
}
