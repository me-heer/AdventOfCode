package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := []string{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
		i++
	}

	part1(matrix)
	part2(matrix)

	println(masCount)

}

var masCount = 0

func part2(matrix []string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			isInsideBoundary := i-1 >= 0 && j-1 >= 0 && j+1 < len(matrix[0]) && i+1 < len(matrix)
			if matrix[i][j] == 'A' && isInsideBoundary {
				// both M's are at the top
				if matrix[i-1][j-1] == 'M' && matrix[i-1][j+1] == 'M' {
					if matrix[i+1][j-1] == 'S' && matrix[i+1][j+1] == 'S' {
						masCount++
					}
				}
				// both M's are at the bottom
				if matrix[i+1][j-1] == 'M' && matrix[i+1][j+1] == 'M' {
					if matrix[i-1][j-1] == 'S' && matrix[i-1][j+1] == 'S' {
						masCount++
					}
				}
				// both M's are at the left
				if matrix[i-1][j-1] == 'M' && matrix[i+1][j-1] == 'M' {
					if matrix[i-1][j+1] == 'S' && matrix[i+1][j+1] == 'S' {
						masCount++
					}
				}
				// both M's are at the right
				if matrix[i-1][j+1] == 'M' && matrix[i+1][j+1] == 'M' {
					if matrix[i-1][j-1] == 'S' && matrix[i+1][j-1] == 'S' {
						masCount++
					}
				}
			}
		}
	}
}

func part1(matrix []string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 'X' {
				// go in all directions
				scanUp(matrix, i, j)
				scanDown(matrix, i, j)
				scanLeft(matrix, i, j)
				scanRight(matrix, i, j)
				scanUpLeft(matrix, i, j)
				scanUpRight(matrix, i, j)
				scanDownLeft(matrix, i, j)
				scanDownRight(matrix, i, j)
			}
		}
	}
}

var xmasCount = 0

func scanLeft(matrix []string, i int, j int) {
	if j-3 < 0 {
		return
	}
	foundX := matrix[i][j] == 'X'
	foundM := matrix[i][j-1] == 'M'
	foundA := matrix[i][j-2] == 'A'
	foundS := matrix[i][j-3] == 'S'
	if foundX && foundM && foundA && foundS {
		xmasCount++
	}
}

func scanRight(matrix []string, i int, j int) {
	if j+3 >= len(matrix[0]) {
		return
	}
	foundX := matrix[i][j] == 'X'
	foundM := matrix[i][j+1] == 'M'
	foundA := matrix[i][j+2] == 'A'
	foundS := matrix[i][j+3] == 'S'
	if foundX && foundM && foundA && foundS {
		xmasCount++
	}
}

func scanUp(matrix []string, i int, j int) {
	if i-3 < 0 {
		return
	}
	foundX := matrix[i][j] == 'X'
	foundM := matrix[i-1][j] == 'M'
	foundA := matrix[i-2][j] == 'A'
	foundS := matrix[i-3][j] == 'S'
	if foundX && foundM && foundA && foundS {
		xmasCount++
	}
}

func scanDown(matrix []string, i int, j int) {
	if i+3 >= len(matrix) {
		return
	}
	foundX := matrix[i][j] == 'X'
	foundM := matrix[i+1][j] == 'M'
	foundA := matrix[i+2][j] == 'A'
	foundS := matrix[i+3][j] == 'S'
	if foundX && foundM && foundA && foundS {
		xmasCount++
	}
}

func scanUpLeft(matrix []string, i int, j int) {
	if i-3 < 0 || j-3 < 0 {
		return
	}
	foundX := matrix[i][j] == 'X'
	foundM := matrix[i-1][j-1] == 'M'
	foundA := matrix[i-2][j-2] == 'A'
	foundS := matrix[i-3][j-3] == 'S'
	if foundX && foundM && foundA && foundS {
		xmasCount++
	}
}

func scanUpRight(matrix []string, i int, j int) {
	if i-3 < 0 || j+3 >= len(matrix[0]) {
		return
	}
	foundX := matrix[i][j] == 'X'
	foundM := matrix[i-1][j+1] == 'M'
	foundA := matrix[i-2][j+2] == 'A'
	foundS := matrix[i-3][j+3] == 'S'
	if foundX && foundM && foundA && foundS {
		xmasCount++
	}
}

func scanDownLeft(matrix []string, i int, j int) {
	if i+3 >= len(matrix) || j-3 < 0 {
		return
	}
	foundX := matrix[i][j] == 'X'
	foundM := matrix[i+1][j-1] == 'M'
	foundA := matrix[i+2][j-2] == 'A'
	foundS := matrix[i+3][j-3] == 'S'
	if foundX && foundM && foundA && foundS {
		xmasCount++
	}
}

func scanDownRight(matrix []string, i int, j int) {
	if i+3 >= len(matrix) || j+3 >= len(matrix[0]) {
		return
	}
	foundX := matrix[i][j] == 'X'
	foundM := matrix[i+1][j+1] == 'M'
	foundA := matrix[i+2][j+2] == 'A'
	foundS := matrix[i+3][j+3] == 'S'
	if foundX && foundM && foundA && foundS {
		xmasCount++
	}
}
