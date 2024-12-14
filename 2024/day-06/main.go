package main

import (
	"bufio"
	"fmt"
	"os"
)

type DirCoordinates struct {
	r      int
	c      int
	dirRow int
	dirCol int
}
type Coordinates struct {
	row int
	col int
}

var visited = make(map[Coordinates]bool)
var reachedBoundary bool
var curr Coordinates

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var matrix []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}

	startX := -1
	startY := -1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '^' {
				startX = i
				startY = j
			}
		}
	}

	curr = Coordinates{row: startX, col: startY}
	visited[curr] = true

	traverse(matrix)
	println(len(visited))

	for visitedCoord, _ := range visited {
		copied := make([]string, len(matrix))
		copy(copied, matrix)
		runes := []rune(copied[visitedCoord.row])
		runes[visitedCoord.col] = '#'
		copied[visitedCoord.row] = string(runes)
		curr = Coordinates{row: startX, col: startY}
		traverseWithLoopDetection(copied)
	}
	println(loop)
}

func traverseWithLoopDetection(matrix []string) {
	history := make(map[DirCoordinates]bool)
	goDirWithLoopDetection(matrix, -1, 0, history)
}

func goDirWithLoopDetection(matrix []string, dirRow int, dirCol int, history map[DirCoordinates]bool) {
	for {
		history[DirCoordinates{r: curr.row, c: curr.col, dirRow: dirRow, dirCol: dirCol}] = true
		isBoundary := curr.row+dirRow < 0 || curr.row+dirRow >= len(matrix) || curr.col+dirCol < 0 || curr.col+dirCol >= len(matrix[0])
		if isBoundary {
			reachedBoundary = true
			return
		}
		isObstacle := matrix[curr.row+dirRow][curr.col+dirCol] == '#'
		if isObstacle {
			if dirRow == -1 && dirCol == 0 {
				dirRow = 0
				dirCol = 1
			} else if dirRow == 0 && dirCol == 1 {
				dirRow = 1
				dirCol = 0
			} else if dirRow == 1 && dirCol == 0 {
				dirRow = 0
				dirCol = -1
			} else if dirRow == 0 && dirCol == -1 {
				dirRow = -1
				dirCol = 0
			}
		} else {
			curr.row = curr.row + dirRow
			curr.col = curr.col + dirCol
		}

		if (history[DirCoordinates{r: curr.row, c: curr.col, dirRow: dirRow, dirCol: dirCol}]) {
			loop++
			return
		}
	}
}

var loop = 0

func traverse(matrix []string) {
	goDir(matrix, -1, 0)
}

func goDir(matrix []string, dirRow int, dirCol int) {
	for {
		isBoundary := curr.row+dirRow < 0 || curr.row+dirRow >= len(matrix) || curr.col+dirCol < 0 || curr.col+dirCol >= len(matrix[0])
		if isBoundary {
			reachedBoundary = true
			return
		}
		isObstacle := matrix[curr.row+dirRow][curr.col+dirCol] == '#'
		if isObstacle {
			if dirRow == -1 && dirCol == 0 {
				dirRow = 0
				dirCol = 1
			} else if dirRow == 0 && dirCol == 1 {
				dirRow = 1
				dirCol = 0
			} else if dirRow == 1 && dirCol == 0 {
				dirRow = 0
				dirCol = -1
			} else if dirRow == 0 && dirCol == -1 {
				dirRow = -1
				dirCol = 0
			}
		}
		curr.row = curr.row + dirRow
		curr.col = curr.col + dirCol
		visited[curr] = true
	}
}
