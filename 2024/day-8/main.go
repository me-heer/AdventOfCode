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
				// antennas[Point{row: i, col: j}] = true
			}
		}
	}

	// for _, p := range antennas {
	// 	println(p.row, " ", p.col)
	// }

	for i := 0; i < len(antennas)-1; i++ {
		for j := 1; j < len(antennas); j++ {
			if i == j {
				continue
			}
			ithAntenna := antennas[i]
			jthAntenna := antennas[j]
			if matrix[ithAntenna.row][ithAntenna.col] == matrix[jthAntenna.row][jthAntenna.col] {
				println("Antenna 1: ", ithAntenna.row, " ", ithAntenna.col)
				println("Antenna 2: ", jthAntenna.row, " ", jthAntenna.col)

				// antinodes possible
				rowDist := (ithAntenna.row - jthAntenna.row)
				if rowDist < 0 {
					rowDist = -rowDist
				}

				colDist := (ithAntenna.col - jthAntenna.col)
				if colDist < 0 {
					colDist = -colDist
				}
				println("ROW DIST: ", rowDist, " COL DIST: ", colDist)

				r1 := 0
				r2 := 0
				if ithAntenna.row <= jthAntenna.row {
					r1 = ithAntenna.row - rowDist
					r2 = jthAntenna.row + rowDist
				} else {
					r1 = jthAntenna.row - rowDist
					r2 = ithAntenna.row + rowDist
				}

				c1 := 0
				c2 := 0
				if ithAntenna.col <= jthAntenna.col {
					c1 = ithAntenna.col - colDist
					c2 = jthAntenna.col + colDist
				} else {
					c1 = jthAntenna.col - colDist
					c2 = ithAntenna.col + colDist
				}

				p1 := Point{}
				p2 := Point{}
				if ithAntenna.row > jthAntenna.row && ithAntenna.col <= jthAntenna.col {
					// i is bottom left
					if r1 > ithAntenna.row {
						p1.row = r1
						p2.row = r2
					} else {
						p1.row = r2
						p2.row = r1
					}

					if c1 <= ithAntenna.col {
						p1.col = c1
						p2.col = c2
					} else {
						p1.col = c2
						p2.col = c1
					}
				} else if ithAntenna.row <= jthAntenna.row && ithAntenna.col > jthAntenna.col {
					// i is top right
					if r1 <= ithAntenna.row {
						p1.row = r1
						p2.row = r2
					} else {
						p1.row = r2
						p2.row = r1
					}

					if c1 > ithAntenna.col {
						p1.col = c1
						p2.col = c2
					} else {
						p1.col = c2
						p2.col = c1
					}
				} else if ithAntenna.row <= jthAntenna.row && ithAntenna.col <= jthAntenna.col {
					// i is top left
					if r1 <= ithAntenna.row {
						p1.row = r1
						p2.row = r2
					} else {
						p1.row = r2
						p2.row = r1
					}

					if c1 <= ithAntenna.col {
						p1.col = c1
						p2.col = c2
					} else {
						p1.col = c2
						p2.col = c1
					}
				} else if ithAntenna.row > jthAntenna.row && ithAntenna.col > jthAntenna.col {
					// i is bottom right
					if r1 > ithAntenna.row {
						p1.row = r1
						p2.row = r2
					} else {
						p1.row = r2
						p2.row = r1
					}

					if c1 > ithAntenna.col {
						p1.col = c1
						p2.col = c2
					} else {
						p1.col = c2
						p2.col = c1
					}
				}

				if p1.row >= 0 && p1.row < len(matrix) && p1.col >= 0 && p1.col < len(matrix[0]) {
					antinodes[p1] = true
				}
				if p2.row >= 0 && p2.row < len(matrix) && p2.col >= 0 && p2.col < len(matrix[0]) {
					antinodes[p2] = true
				}
			}
		}
	}

	for p, _ := range antinodes {
		println(p.row, " ", p.col)
	}

	println(len(antinodes))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if antinodes[Point{row: i, col: j}] {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}
