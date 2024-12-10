package advent

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type step = func(int, int) (int, int)

func up(line, column int) (int, int)        { return line - 1, column }
func down(line, column int) (int, int)      { return line + 1, column }
func left(line, column int) (int, int)      { return line, column - 1 }
func right(line, column int) (int, int)     { return line, column + 1 }
func upRight(line, column int) (int, int)   { return line - 1, column + 1 }
func downRight(line, column int) (int, int) { return line + 1, column + 1 }
func upLeft(line, column int) (int, int)    { return line - 1, column - 1 }
func downLeft(line, column int) (int, int)  { return line + 1, column - 1 }

var stepFunctions = []step{
	upLeft, up, upRight,
	left, right,
	downLeft, down, downRight,
}

const MAS string = "MAS"

type crossWord struct {
	width           int
	heigth          int
	letterToIndexes map[byte][][]int
	chars           [][]byte
}

func (c *crossWord) addLine(line []byte) {
	if c.width == 0 {
		c.width = len(line)
	} else if c.width != len(line) {
		panic(fmt.Sprintf("varying line length in input, expected %d, got %d", c.width, len(line)))
	}
	newLine := make([]byte, len(line))
	copy(newLine, line)

	c.chars = append(c.chars, newLine)
	c.saveLetterPositionsInLine(line, 'X')
	c.saveLetterPositionsInLine(line, 'A')
	c.heigth++
}

func (c *crossWord) saveLetterPositionsInLine(line []byte, letter byte) {
	lineLetterPositions := make([]int, 0, len(line))
	for idx, char := range line {
		if char == letter {
			lineLetterPositions = append(lineLetterPositions, idx)
		}
	}
	if c.letterToIndexes[letter] == nil {
		c.letterToIndexes[letter] = make([][]int, 0, 1028)
	}
	c.letterToIndexes[letter] = append(c.letterToIndexes[letter], lineLetterPositions)
}

func (c crossWord) isOutOfWidthBound(column int) bool {
	return column >= c.width || column < 0
}

func (c crossWord) isOutOfHeigthBound(line int) bool {
	return line >= c.heigth || line < 0
}

func (c crossWord) isWordPresent(stepFn step, word string, line, column int) bool {
	for _, char := range word {
		if c.isOutOfWidthBound(column) {
			return false
		}
		if c.isOutOfHeigthBound(line) {
			return false
		}
		if char != rune(c.chars[line][column]) {
			return false
		}
		line, column = stepFn(line, column)
	}
	return true
}

func (c crossWord) positionXMASCount(line, column int) int {
	sum := 0

	for _, stepFunction := range stepFunctions {
		if c.isWordPresent(stepFunction, "XMAS", line, column) {
			sum++
		}
	}

	return sum
}

func (c crossWord) allXMASCount() int {
	sum := 0
	for lineIdx, line := range c.letterToIndexes['X'] {
		for _, column := range line {
			sum += c.positionXMASCount(lineIdx, column)
		}
	}
	return sum
}

// countMASinXOcurrences count all ocurrences that have the following format:
//
//	M.S
//	.A.
//	M.S
func (c crossWord) countMASinXOcurrences() int {
	sum := 0
	for lineIdx, line := range c.letterToIndexes['A'] {
		if lineIdx > 0 && lineIdx < c.heigth-1 {
			for _, column := range line {
				if c.isMASinX(lineIdx, column) {
					sum++
				}
			}
		}
	}
	return sum
}

// isMASinX checks if an (A) index have MAS in all directions, forward or backward
// which means that it has to appear in 2 out of the possible 4 directions
func (c crossWord) isMASinX(line, column int) bool {
	up := line - 1
	down := line + 1
	left := column - 1
	right := column + 1

	sum := 0
	if c.isWordPresent(downRight, MAS, up, left) {
		sum++
	}
	if c.isWordPresent(downLeft, MAS, up, right) {
		sum++
	}
	if c.isWordPresent(upRight, MAS, down, left) {
		sum++
	}
	if c.isWordPresent(upLeft, MAS, down, right) {
		sum++
	}

	return sum == 2
}

func newCrossWord() crossWord {
	chars := make([][]byte, 0, 256)
	letterToIndexes := make(map[byte][][]int)
	return crossWord{chars: chars, letterToIndexes: letterToIndexes}
}

func newCrossWordFromReader(rd *bufio.Reader) crossWord {
	cross := parseInput(rd)
	return cross
}

func readFile(path string) *bufio.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	return reader
}

func parseInput(reader *bufio.Reader) crossWord {
	cross := newCrossWord()
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		cross.addLine(line)
	}
	return cross
}

func CeresSearchPartOne() {
	path := "advent/4ceressearch/input.txt"
	reader := readFile(path)
	cross := parseInput(reader)
	fmt.Printf("The answer is %d\n", cross.allXMASCount())
}

func CeresSearchPartTwo() {
	path := "advent/4ceressearch/input.txt"
	reader := readFile(path)
	cross := parseInput(reader)
	fmt.Printf("The answer is %d\n", cross.countMASinXOcurrences())
}
