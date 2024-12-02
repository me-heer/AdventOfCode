package main

import (
	"bufio"
	"os"
)

var treesEncountered = 0
var currX = 0
var currY = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var matrix []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		matrix = append(matrix, line)
	}

	// 53, 282, 54, 54, 22
	for {
		move(1, 0, matrix, false)
		move(0, 1, matrix, false)
		move(0, 1, matrix, true)
		if currY == len(matrix)-1 {
			break
		}
	}
	println(treesEncountered)
}

// 282, 53, 54, 54, 22

func move(xDir int, yDir int, matrix []string, check bool) {
	currX = (currX + xDir) % len(matrix[0])
	currY = (currY + yDir) % len(matrix)
	if check && matrix[currY][currX] == '#' {
		treesEncountered++
	}
}
