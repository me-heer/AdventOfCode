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
				fmt.Println("STARTING FOR: ", string(rune(matrix[i][j])))
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
		fmt.Println("CURRENT REGION:", currentRegion[0])
		fmt.Println("CURRENT REGION LEN:", len(currentRegion))
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

		// fmt.Println("AREA: ", area, " PERIMETER: ", perimeter)
		sum += (area * perimeter)
	}
	println(sum)

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
	fmt.Println("CHECKING: ", string(matrix[row][col]), row, col)

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
