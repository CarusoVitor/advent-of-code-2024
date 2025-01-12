package main

import (
	"fmt"

	advent "github.com/CarusoVitor/advent-of-code-2024/advent"
	guardgallivant "github.com/CarusoVitor/advent-of-code-2024/advent/6guardgallivant"
)

func main() {
	path := "advent/6guardgallivant/input.txt"
	rd := advent.ReadFile(path)
	fmt.Println("The answer is:", guardgallivant.GuardGallivantPartTwo(rd))
}
