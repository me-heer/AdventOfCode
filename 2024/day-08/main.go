package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Point struct {
	row int
	col int
}

var antinodes = make(map[Point]bool)
var antennas []Point

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

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			r := rune(matrix[i][j])
			if unicode.IsDigit(r) || unicode.IsUpper(r) || unicode.IsLower(r) {
				antennas = append(antennas, Point{row: i, col: j})
			}
		}
	}

	for i := 0; i < len(antennas)-1; i++ {
		for j := 1; j < len(antennas); j++ {
			if i == j {
				continue
			}
			ithAntenna := antennas[i]
			jthAntenna := antennas[j]
			if matrix[ithAntenna.row][ithAntenna.col] == matrix[jthAntenna.row][jthAntenna.col] {

				for i := 0; i < len(matrix); i++ {
					for j := 0; j < len(matrix[0]); j++ {
						p := Point{row: i, col: j}
						if isColinear(ithAntenna, jthAntenna, p) {
							antinodes[p] = true
						}
					}
				}
			}
		}
	}
	println(len(antinodes))
}

func isColinear(p1 Point, p2 Point, p Point) bool {
	isColinear := (p.row-p1.row)*(p2.col-p1.col) == (p2.row-p1.row)*(p.col-p1.col)
	return isColinear
}
