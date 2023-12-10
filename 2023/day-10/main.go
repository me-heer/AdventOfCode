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
	"S": {},
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

	// find dots,
	// for each dot,
	// go in all directions, keep track of the boundaries, all points in boundaries must be in mainLoop
	var dots []point

	for rowIndex := 0; rowIndex < len(tiles); rowIndex++ {
		for colIndex := 0; colIndex < len(tiles[0]); colIndex++ {
			if !slices.Contains(mainLoop, point{rowIndex, colIndex}) && string(tiles[rowIndex][colIndex]) == "." {
				dots = append(dots, point{rowIndex, colIndex})
			}
		}
		println()
	}

	sum := 0
	for _, p := range dots {
		var boundaries = make(map[point]bool, 0)
		var history = make([]point, 0)
		boundaries, history = floodFill(tiles, p.rowIndex, p.colIndex, boundaries, history)

		areBoundariesValid := true
		if boundaries[point{-1, -1}] {
			areBoundariesValid = false
		}
		if !areBoundariesValid {
			continue
		}

		for point := range boundaries {
			if !slices.Contains(mainLoop, point) {
				areBoundariesValid = false
				break
			}
		}
		sum += len(history)
		println(len(history))
	}
	println(sum)

}

var scanned []point

func isBoundary(s string) bool {
	for boundary := range connectingPaths {
		if s == boundary {
			return true
		}
	}
	return false
}

func floodFill(tiles []string, row, col int, boundaries map[point]bool, history []point) (map[point]bool, []point) {
	if slices.Contains(scanned, point{row, col}) {
		return boundaries, history
	}

	history = append(history, point{row, col})
	scanned = append(scanned, point{row, col})

	if col > 0 {
		next := tiles[row][col-1]
		if isBoundary(string(next)) {
			boundaries[point{row, col - 1}] = true
		} else if !slices.Contains(scanned, point{row, col - 1}) {
			boundaries, history = floodFill(tiles, row, col-1, boundaries, history)
		}
	} else {
		boundaries[point{-1, -1}] = true
	}

	if col < len(tiles[0])-1 {
		next := tiles[row][col+1]
		if isBoundary(string(next)) {
			boundaries[point{row, col + 1}] = true
		} else if !slices.Contains(scanned, point{row, col + 1}) {
			boundaries, history = floodFill(tiles, row, col+1, boundaries, history)
		}
	} else {
		boundaries[point{-1, -1}] = true
	}

	if row > 0 {
		next := tiles[row-1][col]
		if isBoundary(string(next)) {
			boundaries[point{row - 1, col}] = true
		} else if !slices.Contains(scanned, point{row - 1, col}) {
			boundaries, history = floodFill(tiles, row-1, col, boundaries, history)
		}
	} else {
		boundaries[point{-1, -1}] = true
	}

	if row < len(tiles)-1 {
		next := tiles[row+1][col]
		if isBoundary(string(next)) {
			boundaries[point{row + 1, col}] = true
		} else if !slices.Contains(scanned, point{row + 1, col}) {
			boundaries, history = floodFill(tiles, row+1, col, boundaries, history)
		}
	} else {
		boundaries[point{-1, -1}] = true
	}

	return boundaries, history
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
