package main

import (
	"fmt"

	advent "github.com/CarusoVitor/advent-of-code-2024/advent"
	printqueue "github.com/CarusoVitor/advent-of-code-2024/advent/5printqueue"
)

func main() {
	path := "advent/5printqueue/input.txt"
	rd := advent.ReadFile(path)
	fmt.Println("The answer is:", printqueue.PrintQueuePartTwo(rd))
}
