package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	row int
	col int
}

var pointsCovered = make(map[Point]bool)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix []string
	for scanner.Scan() {
		row := scanner.Text()
		matrix = append(matrix, row)
	}

	var allRegions [][]Point
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if !pointsCovered[Point{i, j}] {
				var regionPoints []Point
				regionPoints = expand(matrix, i, j, regionPoints)
				allRegions = append(allRegions, regionPoints)
			}
		}
	}

	sum := 0
	for i := 0; i < len(allRegions); i++ {
		currentRegion := allRegions[i]
		area := len(currentRegion)
		perimeter := 0
		for j := 0; j < len(currentRegion); j++ {
			p := currentRegion[j]

			nr := p.row - 1
			nc := p.col
			newPoint := Point{nr, nc}
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if !contains(currentRegion, newPoint) {
					perimeter++
				}
			} else {
				perimeter++
			}

			nr = p.row + 1
			nc = p.col
			newPoint = Point{nr, nc}
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if !contains(currentRegion, newPoint) {
					perimeter++
				}
			} else {
				perimeter++
			}

			nr = p.row
			nc = p.col - 1
			newPoint = Point{nr, nc}
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if !contains(currentRegion, newPoint) {
					perimeter++
				}
			} else {
				perimeter++
			}

			nr = p.row
			nc = p.col + 1
			newPoint = Point{nr, nc}
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if !contains(currentRegion, newPoint) {
					perimeter++
				}
			} else {
				perimeter++
			}
		}

		sum += (area * perimeter)
	}
	println("PART 1: ", sum)

	var expandedMatrix []string
	for i := 0; i < len(matrix); i++ {
		rowOne := ""
		rowTwo := ""
		for j := 0; j < len(matrix[0]); j++ {
			rowOne += string(matrix[i][j])
			rowOne += string(matrix[i][j])
			rowTwo += string(matrix[i][j])
			rowTwo += string(matrix[i][j])
		}
		expandedMatrix = append(expandedMatrix, rowOne)
		expandedMatrix = append(expandedMatrix, rowTwo)
	}

	println("EXPANDED:")
	for i := 0; i < len(expandedMatrix); i++ {
		for j := 0; j < len(expandedMatrix[0]); j++ {
			fmt.Print(string(expandedMatrix[i][j]))
		}
		println()
	}

	//Clear previous state
	for p := range pointsCovered {
		delete(pointsCovered, p)
	}

	var allRegionsExpanded [][]Point
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if !pointsCovered[Point{i, j}] {
				var regionPoints []Point
				regionPoints = expand(matrix, i, j, regionPoints)
				allRegionsExpanded = append(allRegionsExpanded, regionPoints)
			}
		}
	}

	for i := 0; i < len(allRegionsExpanded); i++ {
		currentRegion := allRegionsExpanded[i]
		area := len(currentRegion)
		corners := 0
		for j := 0; j < len(currentRegion); j++ {
			p := currentRegion[j]
			println("FOR: ", string(expandedMatrix[p.row][p.col]))
			myChar := expandedMatrix[p.row][p.col]

			sameCharSides := make(map[string]bool)
			nr := p.row - 1
			nc := p.col
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if expandedMatrix[nr][nc] == myChar {
					sameCharSides["UP"] = true
				}
			}

			nr = p.row + 1
			nc = p.col
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if expandedMatrix[nr][nc] == myChar {
					sameCharSides["DOWN"] = true
				}
			}

			nr = p.row
			nc = p.col - 1
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if expandedMatrix[nr][nc] == myChar {
					sameCharSides["LEFT"] = true
				}
			}

			nr = p.row
			nc = p.col + 1
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if expandedMatrix[nr][nc] == myChar {
					sameCharSides["RIGHT"] = true
				}
			}

			if len(sameCharSides) == 2 {
				upLeft := sameCharSides["UP"] && sameCharSides["LEFT"]
				upRight := sameCharSides["UP"] && sameCharSides["RIGHT"]
				leftDown := sameCharSides["LEFT"] && sameCharSides["DOWN"]
				downRight := sameCharSides["RIGHT"] && sameCharSides["DOWN"]
				if upLeft || upRight || leftDown || downRight {
					corners++
				}
			}
		}
		println("AREA: ", area)
		println("CORNERS: ", corners)
	}
}

func contains(slice []Point, value Point) bool {
	for _, v := range slice {
		if v.row == value.row && v.col == value.col {
			return true
		}
	}
	return false
}

func expand(matrix []string, row int, col int, regionPoints []Point) []Point {
	if !pointsCovered[Point{row, col}] {
		pointsCovered[Point{row, col}] = true
	} else {
		return regionPoints
	}

	regionPoints = append(regionPoints, Point{row, col})

	nr := row - 1
	nc := col
	if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
		if matrix[nr][nc] == matrix[row][col] {
			regionPoints = expand(matrix, nr, nc, regionPoints)
		}
	}
	nr = row + 1
	nc = col
	if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
		if matrix[nr][nc] == matrix[row][col] {
			regionPoints = expand(matrix, nr, nc, regionPoints)
		}
	}
	nr = row
	nc = col - 1
	if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
		if matrix[nr][nc] == matrix[row][col] {
			regionPoints = expand(matrix, nr, nc, regionPoints)
		}
	}
	nr = row
	nc = col + 1
	if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
		if matrix[nr][nc] == matrix[row][col] {
			regionPoints = expand(matrix, nr, nc, regionPoints)
		}
	}

	return regionPoints
}
