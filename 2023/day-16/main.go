package main

import (
	"bufio"
	"os"
)

type VisitHistory struct {
	point       Point
	visitedFrom string
}

type Point struct {
	x int
	y int
}

const (
	FROM_UP    = "FROM_UP"
	FROM_DOWN  = "FROM_DOWN"
	FROM_LEFT  = "FROM_LEFT"
	FROM_RIGHT = "FROM_RIGHT"
)

const (
	UP    = "UP"
	DOWN  = "DOWN"
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
)

const (
	EMPTY_SPACE         = "."
	SPLITTER_VERTICAL   = "|"
	SPLITTER_HORIZONTAL = "-"
	MIRROR_TILTED_LEFT  = "\\"
	MIRROR_TILTED_RIGHT = "/"
)

func main() {
	input, _ := os.Open("day-16/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var matrix []string

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	max := -1
	for i := 0; i < len(matrix); i++ {
		startRow, startCol := i, -1
		visited := make(map[VisitHistory]bool)
		visited = move(startRow, startCol, matrix, RIGHT, FROM_LEFT, visited)

		var points = make(map[Point]bool)
		for visitHistory := range visited {
			points[visitHistory.point] = true
		}
		result := len(points) - 1
		if result > max {
			max = result
		}
	}

	for i := 0; i < len(matrix); i++ {
		startRow, startCol := i, len(matrix[0])
		visited := make(map[VisitHistory]bool)
		visited = move(startRow, startCol, matrix, LEFT, FROM_RIGHT, visited)

		var points = make(map[Point]bool)
		for visitHistory := range visited {
			points[visitHistory.point] = true
		}
		result := len(points) - 1
		if result > max {
			max = result
		}
	}

	for j := 0; j < len(matrix[0]); j++ {
		startRow, startCol := -1, j
		visited := make(map[VisitHistory]bool)
		visited = move(startRow, startCol, matrix, DOWN, FROM_UP, visited)

		var points = make(map[Point]bool)
		for visitHistory := range visited {
			points[visitHistory.point] = true
		}
		result := len(points) - 1
		if result > max {
			max = result
		}
	}

	for j := 0; j < len(matrix[0]); j++ {
		startRow, startCol := len(matrix), j
		visited := make(map[VisitHistory]bool)
		visited = move(startRow, startCol, matrix, UP, FROM_DOWN, visited)

		var points = make(map[Point]bool)
		for visitHistory := range visited {
			points[visitHistory.point] = true
		}
		result := len(points) - 1
		if result > max {
			max = result
		}
	}

	println(max)
}

func move(currRow, currCol int, matrix []string, direction string, fromDirection string, visited map[VisitHistory]bool) map[VisitHistory]bool {
	if _, ok := visited[VisitHistory{Point{currRow, currCol}, fromDirection}]; ok {
		return visited
	}
	visited[VisitHistory{Point{currRow, currCol}, fromDirection}] = true
	switch direction {
	case UP:
		if currRow > 0 {
			currRow = currRow - 1
		} else {
			return visited
		}
	case DOWN:
		if currRow < len(matrix)-1 {
			currRow = currRow + 1
		} else {
			return visited
		}
	case LEFT:
		if currCol > 0 {
			currCol = currCol - 1
		} else {
			return visited
		}
	case RIGHT:
		if currCol < len(matrix[0])-1 {
			currCol = currCol + 1
		} else {
			return visited
		}
	}
	// moved
	// check what do we have now
	currEncounter := string(matrix[currRow][currCol])
	// EMPTY SPACE: . MIRRORS: / \ SPLITTERS: | -
	switch currEncounter {
	case EMPTY_SPACE:
		// go in the same direction that you were going, from where you were coming
		visited = move(currRow, currCol, matrix, direction, fromDirection, visited)
	case MIRROR_TILTED_LEFT: // \
		switch fromDirection {
		case FROM_UP:
			visited = move(currRow, currCol, matrix, RIGHT, FROM_LEFT, visited)
		case FROM_RIGHT:
			visited = move(currRow, currCol, matrix, UP, FROM_DOWN, visited)
		case FROM_DOWN:
			visited = move(currRow, currCol, matrix, LEFT, FROM_RIGHT, visited)
		case FROM_LEFT:
			visited = move(currRow, currCol, matrix, DOWN, FROM_UP, visited)
		}
	case MIRROR_TILTED_RIGHT: // /
		switch fromDirection {
		case FROM_UP:
			visited = move(currRow, currCol, matrix, LEFT, FROM_RIGHT, visited)
		case FROM_RIGHT:
			visited = move(currRow, currCol, matrix, DOWN, FROM_UP, visited)
		case FROM_DOWN:
			visited = move(currRow, currCol, matrix, RIGHT, FROM_LEFT, visited)
		case FROM_LEFT:
			visited = move(currRow, currCol, matrix, UP, FROM_DOWN, visited)
		}
	case SPLITTER_HORIZONTAL: // -
		switch fromDirection {
		case FROM_UP, FROM_DOWN:
			visited = move(currRow, currCol, matrix, LEFT, FROM_RIGHT, visited)
			visited = move(currRow, currCol, matrix, RIGHT, FROM_LEFT, visited)
		case FROM_LEFT, FROM_RIGHT:
			visited = move(currRow, currCol, matrix, direction, fromDirection, visited)
		}
	case SPLITTER_VERTICAL: // |
		switch fromDirection {
		case FROM_UP, FROM_DOWN:
			visited = move(currRow, currCol, matrix, direction, fromDirection, visited)
		case FROM_LEFT, FROM_RIGHT:
			visited = move(currRow, currCol, matrix, UP, FROM_DOWN, visited)
			visited = move(currRow, currCol, matrix, DOWN, FROM_UP, visited)
		}
	}

	return visited
}
