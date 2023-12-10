package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"slices"
	"strings"
)

type point struct {
	rowIndex int
	colIndex int
}

var mainLoop []point

var minDistance = make(map[point]int)

var connectingPaths = map[string]map[string]func(int, int) (int, int, string){
	// if coming from down, go right OR if coming from right, go down
	"F": {
		"fromdown": func(row, col int) (int, int, string) {
			col++ // go right
			return row, col, "fromleft"
		},
		"fromright": func(row, col int) (int, int, string) {
			row++ // go down
			return row, col, "fromup"
		},
	},
	"|": {
		"fromup": func(row, col int) (int, int, string) {
			row++ // go down
			return row, col, "fromup"
		},
		"fromdown": func(row, col int) (int, int, string) {
			row-- // go up
			return row, col, "fromdown"
		},
	},
	"-": {
		"fromleft": func(row, col int) (int, int, string) {
			col++ // go right
			return row, col, "fromleft"
		},
		"fromright": func(row, col int) (int, int, string) {
			col-- // go left
			return row, col, "fromright"
		},
	},
	"L": {
		"fromup": func(row, col int) (int, int, string) {
			col++ // go right
			return row, col, "fromleft"
		},
		"fromright": func(row, col int) (int, int, string) {
			row-- // go up
			return row, col, "fromdown"
		},
	},
	"J": {
		"fromup": func(row, col int) (int, int, string) {
			col-- // go left
			return row, col, "fromright"
		},
		"fromleft": func(row, col int) (int, int, string) {
			row-- //go up
			return row, col, "fromdown"
		},
	},
	"7": {
		"fromdown": func(row, col int) (int, int, string) {
			col-- //goleft
			return row, col, "fromright"
		},
		"fromleft": func(row, col int) (int, int, string) {
			row++ //go down
			return row, col, "fromup"
		},
	},
}

var tiles []string

func main() {
	input, err := os.Open("day-10/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = strings.TrimRight(line, "\n")
		tiles = append(tiles, line)
	}

	// find Start
	var startRow, startCol int
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles[0]); j++ {
			if string(tiles[i][j]) == "S" {
				startRow = i
				startCol = j
				break
			}
		}
	}

	//scan around and find connecting paths
	if startRow-1 > 0 {
		if paths, ok := connectingPaths[string(tiles[startRow-1][startCol])]; ok {
			if _, ok := paths["fromdown"]; ok {
				walk(startRow-1, startCol, string(tiles[startRow-1][startCol]), "fromdown", 1, []point{{startRow - 1, startCol}})
			}
		}

	}
	if paths, ok := connectingPaths[string(tiles[startRow+1][startCol])]; ok {
		if _, ok := paths["fromup"]; ok {
			walk(startRow+1, startCol, string(tiles[startRow+1][startCol]), "fromup", 1, []point{{startRow + 1, startCol}})
		}
	}

	if startCol-1 > 0 {
		if paths, ok := connectingPaths[string(tiles[startRow][startCol-1])]; ok {
			if _, ok := paths["fromright"]; ok {
				walk(startRow, startCol-1, string(tiles[startRow][startCol-1]), "fromright", 1, []point{{startRow, startCol - 1}})
			}
		}
	}

	if paths, ok := connectingPaths[string(tiles[startRow][startCol+1])]; ok {
		if _, ok := paths["fromleft"]; ok {
			walk(startRow, startCol+1, string(tiles[startRow][startCol+1]), "fromleft", 1, []point{{startRow, startCol + 1}})
		}
	}

	farthestPointDistance := math.MinInt
	for _, distanceToPoint := range minDistance {
		if distanceToPoint > farthestPointDistance {
			farthestPointDistance = distanceToPoint
		}
	}

	// part2
	// print input map
	println("INPUT MAP:")
	for rowIndex := 0; rowIndex < len(tiles); rowIndex++ {
		for colIndex := 0; colIndex < len(tiles[0]); colIndex++ {
			if slices.Contains(mainLoop, point{rowIndex, colIndex}) {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}

	println("INPUT MAP INVERTED:")
	for rowIndex := 0; rowIndex < len(tiles); rowIndex++ {
		for colIndex := 0; colIndex < len(tiles[0]); colIndex++ {
			if slices.Contains(mainLoop, point{rowIndex, colIndex}) {
				print(".")
			} else {
				print("#")
			}
		}
		println()
	}

	// horizontal scan
	horizontallyScannedTiles := make([][]string, len(tiles))
	for i := 0; i < len(tiles); i++ {
		horizontallyScannedTiles[i] = make([]string, len(tiles[0]))
	}

	for rowIndex := 0; rowIndex < len(tiles); rowIndex++ {
		// find leftMost mainLoop point, rightMost mainLoop point
		leftMostMainLoopPointColIndex := math.MaxInt
		rightMostMainLoopPointColIndex := math.MinInt
		for _, pointInMainLoop := range mainLoop {
			if pointInMainLoop.rowIndex == rowIndex {
				if pointInMainLoop.colIndex < leftMostMainLoopPointColIndex {
					leftMostMainLoopPointColIndex = pointInMainLoop.colIndex
				}
				if pointInMainLoop.colIndex > rightMostMainLoopPointColIndex {
					rightMostMainLoopPointColIndex = pointInMainLoop.colIndex
				}
			}
		}

		for colIndex := 0; colIndex < len(tiles[0]); colIndex++ {
			if slices.Contains(mainLoop, point{rowIndex, colIndex}) {
				horizontallyScannedTiles[rowIndex][colIndex] = "#"
			} else if colIndex >= leftMostMainLoopPointColIndex && colIndex <= rightMostMainLoopPointColIndex {
				horizontallyScannedTiles[rowIndex][colIndex] = "I"
			} else {
				horizontallyScannedTiles[rowIndex][colIndex] = "."
			}
		}
	}

	// vertical scan
	verticallyScannedTiles := make([][]string, len(tiles))
	for i := 0; i < len(tiles); i++ {
		verticallyScannedTiles[i] = make([]string, len(tiles[0]))
	}

	for colIndex := 0; colIndex < len(tiles[0]); colIndex++ {
		// find topMost mainLoop point, bottomMost mainLoop point
		topMostMainLoopRowIndex := math.MaxInt
		bottomMostMainLoopRowIndex := math.MinInt
		for _, pointInMainLoop := range mainLoop {
			if pointInMainLoop.colIndex == colIndex {
				if pointInMainLoop.rowIndex < topMostMainLoopRowIndex {
					topMostMainLoopRowIndex = pointInMainLoop.rowIndex
				}
				if pointInMainLoop.rowIndex > bottomMostMainLoopRowIndex {
					bottomMostMainLoopRowIndex = pointInMainLoop.rowIndex
				}
			}
		}

		for rowIndex := 0; rowIndex < len(tiles); rowIndex++ {
			if slices.Contains(mainLoop, point{rowIndex, colIndex}) {
				verticallyScannedTiles[rowIndex][colIndex] = "#"
			} else if rowIndex >= topMostMainLoopRowIndex && topMostMainLoopRowIndex <= bottomMostMainLoopRowIndex {
				verticallyScannedTiles[rowIndex][colIndex] = "I"
			} else {
				verticallyScannedTiles[rowIndex][colIndex] = "."
			}
		}
	}

	println("WE DID IT BRO")
	for rowIndex := 0; rowIndex < len(tiles); rowIndex++ {
		for colIndex := 0; colIndex < len(tiles[0]); colIndex++ {
			if horizontallyScannedTiles[rowIndex][colIndex] == "#" {
				print("#")
			} else if horizontallyScannedTiles[rowIndex][colIndex] == "I" && verticallyScannedTiles[rowIndex][colIndex] == "I" {
				print("I")
			} else {
				print(".")
			}
		}
		println()
	}

}

func walk(row, col int, connectingPath string, fromDir string, totalSteps int, history []point) {
	previousMinDistance, ok := minDistance[point{row, col}]
	if !ok {
		minDistance[point{row, col}] = totalSteps
	} else if ok && totalSteps < previousMinDistance {
		minDistance[point{row, col}] = totalSteps
	}
	moveInDirection, ok := connectingPaths[connectingPath][fromDir]
	if ok {
		row, col, fromDir = moveInDirection(row, col)
		newCurrentTile := string(tiles[row][col])
		history = append(history, point{row, col})
		walk(row, col, newCurrentTile, fromDir, totalSteps+1, history)
	}
	if connectingPath == "S" {
		minDistance[point{row, col}] = 0
		mainLoop = history
	}
}
