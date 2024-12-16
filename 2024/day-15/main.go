package main

import (
	"bufio"
	"fmt"
	"os"
)

var matrix [][]int
var largeMatrix [][]int
var currRow, currCol int
var WALL = -1
var ROBOT = 1
var BOX = 2
var BIGGER_BOX_LEFT = 3
var BIGGER_BOX_RIGHT = 4
var EMPTY_SPACE = 0

type Point struct {
	row int
	col int
}

type Box struct {
	left  Point
	right Point
}

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
		var row []int
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == WALL {
				row = append(row, WALL)
				row = append(row, WALL)
			}
			if matrix[i][j] == BOX {
				row = append(row, BIGGER_BOX_LEFT)
				row = append(row, BIGGER_BOX_RIGHT)
			}
			if matrix[i][j] == EMPTY_SPACE {
				row = append(row, EMPTY_SPACE)
				row = append(row, EMPTY_SPACE)
			}
			if matrix[i][j] == ROBOT {
				row = append(row, ROBOT)
				row = append(row, EMPTY_SPACE)
			}
		}
		largeMatrix = append(largeMatrix, row)
	}

	matrix = largeMatrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == ROBOT {
				currRow = i
				currCol = j
			}
		}
	}

	for i := 0; i < len(largeMatrix); i++ {
		for j := 0; j < len(largeMatrix[0]); j++ {
			if largeMatrix[i][j] == WALL {
				print("#")
			}
			if largeMatrix[i][j] == ROBOT {
				print("@")
			}
			if largeMatrix[i][j] == EMPTY_SPACE {
				print(".")
			}
			if largeMatrix[i][j] == BIGGER_BOX_LEFT {
				print("[")
			}
			if largeMatrix[i][j] == BIGGER_BOX_RIGHT {
				print("]")
			}
		}
		println()
	}

	for _, move := range moves {
		for _, m := range move {
			moveDir(m)
		}
	}

	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == BIGGER_BOX_LEFT {
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
	println("DIRROW: ", dirRow, " DIRCOL: ", dirCol)

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

	case BIGGER_BOX_LEFT, BIGGER_BOX_RIGHT:
		println("NEXT IS BIGGER_BOX")
		// 1. space after box is empty -> move
		// 2. space after box is another box (1 or more boxes) -> capture all the indexes of such boxes
		//    and then check what's next -> next could be . or #
		//    if # do nothing, if . then move all boxes +1

		var nextPoints []Point
		nextPoints = append(nextPoints, Point{currRow + dirRow, currCol + dirCol})

		// var currBox Box

		// if nextPos == BIGGER_BOX_LEFT {
		// 	p := Point{currRow + dirRow, currCol + dirCol}
		// 	currBox.left = p
		// 	p.col++
		// 	currBox.right = p
		// 	nextPoints = append(nextPoints, p)
		// } else if nextPos == BIGGER_BOX_RIGHT {
		// 	p := Point{currRow + dirRow, currCol + dirCol}
		// 	currBox.right = p
		// 	p.col--
		// 	currBox.left = p
		// 	nextPoints = append(nextPoints, p)
		// }

		pointsToBeScanned := make([]Point, len(nextPoints))
		copy(pointsToBeScanned, nextPoints)

		var consecutiveBoxes []Box
		// consecutiveBoxes = append(consecutiveBoxes, currBox)

		for _, p := range pointsToBeScanned {
			p.row += dirRow
			p.col += dirCol
		}

		foundWall := false

		println("CURR: ", currRow, currCol)
		println("POINTS TO BE SCANNED...:")
		for _, p := range pointsToBeScanned {
			println(p.row, p.col)
		}

		for len(pointsToBeScanned) > 0 {
			var nextPointsToBeScanned []Point

			for _, p := range pointsToBeScanned {
				if !isInBounds(p.row, p.col) {
					continue
				}
				if matrix[p.row][p.col] == WALL {
					foundWall = true
					break
				}
				if matrix[p.row][p.col] == EMPTY_SPACE {
				}
				if matrix[p.row][p.col] == BIGGER_BOX_LEFT {
					boxLeft := p
					boxRight := p
					boxRight.col++
					newBox := Box{boxLeft, boxRight}
					consecutiveBoxes = append(consecutiveBoxes, newBox)
					leftPoint := Point{newBox.left.row + dirRow, newBox.left.col + dirCol}
					rightPoint := Point{newBox.right.row + dirRow, newBox.right.col + dirCol}
					if leftPoint != boxRight {
						nextPointsToBeScanned = append(nextPointsToBeScanned, leftPoint)
					}

					if rightPoint != boxLeft {
						nextPointsToBeScanned = append(nextPointsToBeScanned, rightPoint)
					}
				}
				if matrix[p.row][p.col] == BIGGER_BOX_RIGHT {
					boxRight := p
					boxLeft := p
					boxLeft.col--
					newBox := Box{boxLeft, boxRight}
					consecutiveBoxes = append(consecutiveBoxes, newBox)
					leftPoint := Point{newBox.left.row + dirRow, newBox.left.col + dirCol}
					rightPoint := Point{newBox.right.row + dirRow, newBox.right.col + dirCol}
					if leftPoint != boxRight {
						nextPointsToBeScanned = append(nextPointsToBeScanned, leftPoint)
					}

					if rightPoint != boxLeft {
						nextPointsToBeScanned = append(nextPointsToBeScanned, rightPoint)
					}
				}
			}

			pointsToBeScanned = nextPointsToBeScanned
			println("NEXT POINTS TO BE SCANNED...:")
			for _, p := range pointsToBeScanned {
				println(p.row, p.col)
			}
		}

		println("CONSECUTIVE BOXES: ", len(consecutiveBoxes))
		var updatedConsecutiveBoxes []Box

		for i := 0; i < len(consecutiveBoxes); i++ {
			updatedBox := consecutiveBoxes[i]
			updatedBox.left = Point{updatedBox.left.row + dirRow, updatedBox.left.col + dirCol}
			updatedBox.right = Point{updatedBox.right.row + dirRow, updatedBox.right.col + dirCol}
			updatedConsecutiveBoxes = append(updatedConsecutiveBoxes, updatedBox)
		}

		if foundWall {
			// can't do anything
		} else {
			// move all consecutive boxes
			for _, b := range consecutiveBoxes {
				matrix[b.left.row][b.left.col] = EMPTY_SPACE
				matrix[b.right.row][b.right.col] = EMPTY_SPACE
			}
			for _, b := range updatedConsecutiveBoxes {
				matrix[b.left.row][b.left.col] = BIGGER_BOX_LEFT
				matrix[b.right.row][b.right.col] = BIGGER_BOX_RIGHT
			}

			matrix[currRow][currCol] = EMPTY_SPACE
			matrix[currRow+dirRow][currCol+dirCol] = ROBOT
			currRow += dirRow
			currCol += dirCol
		}

	case WALL:

	}

	println("NOW CURR: ", currRow, currCol)
	println("MOVED: ", string(dir))
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
	// 		if matrix[i][j] == BIGGER_BOX_LEFT {
	// 			print("[")
	// 		}
	// 		if matrix[i][j] == BIGGER_BOX_RIGHT {
	// 			print("]")
	// 		}
	// 	}
	// 	println()
	// }

}

func isInBounds(currRow int, currCol int) bool {
	return currRow >= 0 && currRow < len(matrix) && currCol >= 0 && currCol < len(matrix[0])
}
