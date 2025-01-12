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

type movingObject struct {
	pos      position
	pointing direction
}

func newMovingObject(column, row int, pointing direction) movingObject {
	pos := newPosition(column, row)
	return movingObject{pos: pos, pointing: pointing}

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
	pos         position
	pointing    direction
	startingPos position
}

func (g *guard) setPosition(column, row int) {
	pos := newPosition(column, row)
	g.pos = pos
}

func (g *guard) setStartingPosition(column, row int) {
	pos := newPosition(column, row)
	g.startingPos = pos
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

func (g *grid) setGuardStartingPosition(column, row int) {
	g.guardObj.setStartingPosition(column, row)
}

func (g grid) isGuardInBounds() bool {
	isRowValid := g.guardObj.row() >= 0 && g.guardObj.row() < g.width
	isColumnValid := g.guardObj.column() >= 0 && g.guardObj.column() < g.heigth
	return isRowValid && isColumnValid
}

func (g grid) at(column, row int) byte {
	return g.lines[row][column]
}

func (g *grid) placeObstacle(column, row int) {
	g.lines[row][column] = '#'
}

func (g *grid) removeObstacle(column, row int) {
	g.lines[row][column] = '.'
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
				grid.setGuardStartingPosition(column, row)
				grid.setGuardPosition(column, row)
			}
			line = append(line, char)
			column++
		}
	}
	return grid
}

func countPositionsUntilLeave(area grid) (int, map[position]bool, bool) {
	count := 0
	visited := make(map[position]bool)
	directionsHistory := make(map[movingObject]bool)

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
			posWithDirection := newMovingObject(column, row, area.guardObj.pointing)
			// if we have already visited this position with this direction, we are in a loop
			if directionsHistory[posWithDirection] {
				return count, visited, true
			}
			directionsHistory[posWithDirection] = true
		}
		area.guardObj.step()
	}
	return count, visited, false
}

func GuardGallivantPartOne(rd *bufio.Reader) int {
	grid := parse(rd)
	count, _, _ := countPositionsUntilLeave(grid)
	return count
}

func GuardGallivantPartTwo(rd *bufio.Reader) int {
	grid := parse(rd)
	_, visited, _ := countPositionsUntilLeave(grid)
	count := countPlacingsThatMakeLoops(grid, visited)
	return count
}

func countPlacingsThatMakeLoops(area grid, visited map[position]bool) int {
	count := 0
	for pos := range visited {
		if pos == area.guardObj.startingPos {
			continue
		}
		area.placeObstacle(pos.column, pos.row)
		_, _, isLoop := countPositionsUntilLeave(area)
		if isLoop {
			count++
		}
		area.removeObstacle(pos.column, pos.row)
	}
	return count
}
