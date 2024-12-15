package main

import (
	"bufio"
	"fmt"
	"os"
)

var matrix [][]int
var currRow, currCol int
var WALL = -1
var ROBOT = 1
var BOX = 2
var EMPTY_SPACE = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var row []int
		for _, r := range line {
			if r == '#' {
				row = append(row, -1)
			}
			if r == '@' {
				row = append(row, 1)
			}
			if r == 'O' {
				row = append(row, 2)
			}
			if r == '.' {
				row = append(row, 0)
			}
		}
		matrix = append(matrix, row)
	}

	var moves []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		moves = append(moves, line)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				currRow = i
				currCol = j
			}
		}
	}

	for _, move := range moves {
		for _, m := range move {
			moveDir(m)
		}
	}

	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == BOX {
				sum += (100*i + j)

			}
		}
	}
	println(sum)
}

func moveDir(dir rune) {
	var dirRow, dirCol int
	switch dir {
	case '<':
		dirRow = 0
		dirCol = -1
	case '>':
		dirRow = 0
		dirCol = +1
	case '^':
		dirRow = -1
		dirCol = 0
	case 'v':
		dirRow = +1
		dirCol = 0
	}

	nextPos := matrix[currRow+dirRow][currCol+dirCol]
	switch nextPos {
	case EMPTY_SPACE:
		println("NEXT IS EMPTY SPACE")
		println("PUT EMPTY SPACE AT ", currRow, currCol)
		println("PUT ROBOT AT ", currRow+dirRow, currCol+dirCol)
		matrix[currRow][currCol] = EMPTY_SPACE
		matrix[currRow+dirRow][currCol+dirCol] = ROBOT
		currRow = currRow + dirRow
		currCol = currCol + dirCol
	case BOX:
		println("NEXT IS BOX")
		// 1. space after box is empty -> move
		// 2. space after box is another box (1 or more boxes) -> capture all the indexes of such boxes
		//    and then check what's next -> next could be . or #
		//    if # do nothing, if . then move all boxes +1
		consecutiveBoxes := 0
		boxCurrRow := currRow
		boxCurrCol := currCol

		for isInBounds(boxCurrRow+dirRow, boxCurrCol+dirCol) && matrix[boxCurrRow+dirRow][boxCurrCol+dirCol] == BOX {
			boxCurrRow += dirRow
			boxCurrCol += dirCol
			consecutiveBoxes++
		}
		println("CONSECUTIVE BOXES: ", consecutiveBoxes)

		if isInBounds(boxCurrRow+dirRow, boxCurrCol+dirCol) && matrix[boxCurrRow+dirRow][boxCurrCol+dirCol] == EMPTY_SPACE {
			println("PUT EMPTY SPACE AT", currRow, currCol)
			matrix[currRow][currCol] = EMPTY_SPACE
			matrix[currRow+dirRow][currCol+dirCol] = ROBOT
			matrix[boxCurrRow+dirRow][boxCurrCol+dirCol] = BOX
			currRow += dirRow
			currCol += dirCol
		} else if isInBounds(boxCurrRow+dirRow, boxCurrCol+dirCol) && matrix[boxCurrRow+dirRow][boxCurrCol+dirCol] == '#' {

		}

	case WALL:

	}

	// println("MOVED: ", string(dir))
	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[0]); j++ {
	// 		if matrix[i][j] == WALL {
	// 			print("#")
	// 		}
	// 		if matrix[i][j] == ROBOT {
	// 			print("@")
	// 		}
	// 		if matrix[i][j] == EMPTY_SPACE {
	// 			print(".")
	// 		}
	// 		if matrix[i][j] == BOX {
	// 			print("O")
	// 		}
	// 	}
	// 	println()
	// }

}

func isInBounds(currRow int, currCol int) bool {
	return currRow >= 0 && currRow < len(matrix) && currCol >= 0 && currCol < len(matrix[0])
}

func updateStringAtIndex(str string, index int, newChar rune) (string, error) {
	// Ensure the index is valid
	if index < 0 || index >= len(str) {
		return str, fmt.Errorf("index out of bounds")
	}

	// Convert string to []rune to handle Unicode safely
	runes := []rune(str)

	// Update the character at the specified index
	runes[index] = newChar

	// Convert runes back to a string
	return string(runes), nil
}
