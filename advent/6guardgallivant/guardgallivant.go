package advent

import (
	"bufio"
	"fmt"
	"io"
)

type position struct {
	column int
	row    int
}

func newPosition(column, row int) position {
	return position{column, row}
}

type direction int

const (
	LEFT direction = iota
	RIGHT
	UP
	DOWN
)

const GUARD_STARTING_DIRECTION = UP

type guard struct {
	pos      position
	pointing direction
}

func (g *guard) setPosition(column, row int) {
	pos := newPosition(column, row)
	g.pos = pos
}

func (g *guard) turnRight90Degrees() {
	if g.pointing == UP {
		g.pointing = RIGHT
	} else if g.pointing == RIGHT {
		g.pointing = DOWN
	} else if g.pointing == DOWN {
		g.pointing = LEFT
	} else if g.pointing == LEFT {
		g.pointing = UP
	}
}

func (g *guard) step() {
	if g.pointing == UP {
		g.pos.row--
	} else if g.pointing == RIGHT {
		g.pos.column++
	} else if g.pointing == DOWN {
		g.pos.row++
	} else if g.pointing == LEFT {
		g.pos.column--
	}
}

func (g *guard) undoStep() {
	if g.pointing == UP {
		g.pos.row++
	} else if g.pointing == RIGHT {
		g.pos.column--
	} else if g.pointing == DOWN {
		g.pos.row--
	} else if g.pointing == LEFT {
		g.pos.column++
	}
}

func (g guard) column() int {
	return g.pos.column
}

func (g guard) row() int {
	return g.pos.row
}

func newGuard() guard {
	return guard{pointing: GUARD_STARTING_DIRECTION}
}

type grid struct {
	lines    [][]byte
	guardObj guard
	width    int
	heigth   int
}

func newGrid() grid {
	lines := make([][]byte, 0, 512)
	guard := newGuard()
	return grid{lines: lines, guardObj: guard}
}

const invalidWidthMsg = "all lines should have width %d, but got %d in line %s"

func (g *grid) addLine(line []byte) {
	if g.width == 0 {
		g.width = len(line)
	} else if len(line) != g.width {
		panic(fmt.Sprintf(invalidWidthMsg, g.width, len(line), line))
	}
	g.lines = append(g.lines, line)
	g.heigth++
}

func (g *grid) setGuardPosition(column, row int) {
	g.guardObj.setPosition(column, row)
}

func (g grid) isGuardInBounds() bool {
	isRowValid := g.guardObj.row() >= 0 && g.guardObj.row() < g.width
	isColumnValid := g.guardObj.column() >= 0 && g.guardObj.column() < g.heigth
	return isRowValid && isColumnValid
}

func (g grid) at(column, row int) byte {
	return g.lines[row][column]
}

func parse(rd *bufio.Reader) grid {
	grid := newGrid()
	line := make([]byte, 0, 128)
	row := 0
	column := 0
	eof := false
	for !eof {
		char, err := rd.ReadByte()
		eof = err == io.EOF
		if char == '\n' || eof {
			row++
			column = 0
			grid.addLine(line)
			line = make([]byte, 0, 128)
		} else {
			if char == '^' {
				grid.setGuardPosition(column, row)
			}
			line = append(line, char)
			column++
		}
	}
	return grid
}

func countPositionsUntilLeave(area grid) int {
	count := 0
	visited := make(map[position]bool)
	for area.isGuardInBounds() {
		row := area.guardObj.row()
		column := area.guardObj.column()
		element := area.at(column, row)
		if element == '#' {
			area.guardObj.undoStep()
			area.guardObj.turnRight90Degrees()
		} else {
			pos := newPosition(column, row)
			if !visited[pos] {
				visited[pos] = true
				count++
			}
		}
		area.guardObj.step()
	}
	return count
}

func GuardGallivantPartOne(rd *bufio.Reader) int {
	grid := parse(rd)
	count := countPositionsUntilLeave(grid)
	return count
}
