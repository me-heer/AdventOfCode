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

	result := 0
	for i := 0; i < len(allRegions); i++ {
		currentRegion := allRegions[i]
		area := len(currentRegion)

		var leftSidePoints []Point
		var rightSidePoints []Point
		var upSidePoints []Point
		var downSidePoints []Point
		for j := 0; j < len(currentRegion); j++ {
			p := currentRegion[j]
			myChar := matrix[p.row][p.col]

			nr := p.row
			nc := p.col - 1
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if matrix[nr][nc] != myChar {
					leftSidePoints = append(leftSidePoints, p)
				}
			} else {
				leftSidePoints = append(leftSidePoints, p)
			}

			nr = p.row
			nc = p.col + 1
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if matrix[nr][nc] != myChar {
					rightSidePoints = append(rightSidePoints, p)
				}
			} else {
				rightSidePoints = append(rightSidePoints, p)
			}

			nr = p.row - 1
			nc = p.col
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if matrix[nr][nc] != myChar {
					upSidePoints = append(upSidePoints, p)
				}
			} else {
				upSidePoints = append(upSidePoints, p)
			}

			nr = p.row + 1
			nc = p.col
			if nr >= 0 && nr < len(matrix) && nc >= 0 && nc < len(matrix[0]) {
				if matrix[nr][nc] != myChar {
					downSidePoints = append(downSidePoints, p)
				}
			} else {
				downSidePoints = append(downSidePoints, p)
			}
		}

		leftGroups := groupAdjacentPoints(leftSidePoints, "LEFT")
		rightGroups := groupAdjacentPoints(rightSidePoints, "RIGHT")
		upGroups := groupAdjacentPoints(upSidePoints, "UP")
		downGroups := groupAdjacentPoints(downSidePoints, "DOWN")

		lg := 0
		for _, group := range leftGroups {
			if len(group) > 0 {
				lg++
			}
		}
		rg := 0
		for _, group := range rightGroups {
			if len(group) > 0 {
				rg++
			}
		}
		ug := 0
		for _, group := range upGroups {
			if len(group) > 0 {
				ug++
			}
		}
		dg := 0
		for _, group := range downGroups {
			if len(group) > 0 {
				dg++
			}
		}

		sides := (lg) + (rg) + (ug) + (dg)
		result += area * sides
	}
	println(result)
}

func groupAdjacentPoints(points []Point, forSide string) [][]Point {
	var adjacentMatrix [][]Point
	for i := 0; i < len(points); i++ {
		curr := points[i]
		groupId := -1
		var currGroup []Point
		for g := 0; g < len(adjacentMatrix); g++ {
			if contains(adjacentMatrix[g], curr) {
				groupId = g
				break
			}
		}

		if groupId != -1 {
			currGroup = adjacentMatrix[groupId]
		} else {
			currGroup = append(currGroup, curr)
		}

		var groupToBeMerged = -1
		for j := 0; j < len(points); j++ {
			if j == i {
				continue
			}
			p1 := points[i]
			p2 := points[j]

			isAdjecent := false
			if forSide == "LEFT" || forSide == "RIGHT" {
				diff := p1.row - p2.row
				if diff < 0 {
					diff = -diff
				}
				if diff == 1 && p1.col == p2.col {
					isAdjecent = true
				}
			} else if forSide == "UP" || forSide == "DOWN" {
				diff := p1.col - p2.col
				if diff < 0 {
					diff = -diff
				}
				if diff == 1 && p1.row == p2.row {
					isAdjecent = true
				}
			}

			if isAdjecent {

				adjecentPointGroup := -1
				for g := 0; g < len(adjacentMatrix); g++ {
					if contains(adjacentMatrix[g], points[j]) && !contains(adjacentMatrix[g], curr) {
						groupToBeMerged = g
						adjecentPointGroup = g
						break
					}
				}

				if !contains(currGroup, points[j]) && adjecentPointGroup == -1 {
					currGroup = append(currGroup, points[j])
				} else if !contains(currGroup, points[j]) && adjecentPointGroup != -1 {
					groupId = adjecentPointGroup
				}
			}
		}

		if groupId == -1 {
			adjacentMatrix = append(adjacentMatrix, currGroup)
		} else {
			if groupToBeMerged != -1 {
				before := adjacentMatrix[groupToBeMerged]
				before = append(before, currGroup...)
				adjacentMatrix[groupId] = []Point{}
				adjacentMatrix[groupToBeMerged] = before
			} else {
				adjacentMatrix[groupId] = currGroup
			}

		}
	}
	return adjacentMatrix
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
