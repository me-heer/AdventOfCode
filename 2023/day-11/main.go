package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("day-11/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)

	var universe [][]string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = strings.TrimRight(line, "\n")
		var lineStr = make([]string, len(line))
		for i := range line {
			lineStr[i] = string(line[i])
		}
		universe = append(universe, lineStr)
	}
	emptyGalaxyRows, emptyGalaxyColumns := expand(universe)
	galaxies := findGalaxies(universe)
	sum := float64(0)
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			// distance between i and j galaxy
			result := math.Abs(float64(galaxies[i][0]-galaxies[j][0])) + math.Abs(float64(galaxies[i][1]-galaxies[j][1]))
			x1 := galaxies[i][0]
			x2 := galaxies[j][0]

			if galaxies[i][0] > galaxies[j][0] {
				x1 = galaxies[j][0]
				x2 = galaxies[i][0]
			}

			add := float64(0)
			for _, rowIndex := range emptyGalaxyRows {
				if x1 < rowIndex && rowIndex < x2 {
					add += 999999
				}
			}
			result += add

			y1 := galaxies[i][1]
			y2 := galaxies[j][1]
			if galaxies[i][1] > galaxies[j][1] {
				y1 = galaxies[j][1]
				y2 = galaxies[i][1]
			}
			add = float64(0)
			for _, colIndex := range emptyGalaxyColumns {
				if y1 < colIndex && colIndex < y2 {
					add += 999999
				}
			}
			result += add
			sum += result
		}
	}
	println(int(sum))
}

func findGalaxies(universe [][]string) [][]int {
	var indexesOfGalaxies [][]int
	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[0]); j++ {
			if universe[i][j] == "#" {
				galaxy := []int{i, j}
				indexesOfGalaxies = append(indexesOfGalaxies, galaxy)
			}
		}
	}
	return indexesOfGalaxies
}

func expand(universe [][]string) ([]int, []int) {
	var emptyGalaxyRows []int
	for i := 0; i < len(universe); i++ {
		isEmpty := true
		for j := 0; j < len(universe[i]); j++ {
			if universe[i][j] == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyGalaxyRows = append(emptyGalaxyRows, i)
		}
	}

	var emptyGalaxyColumns []int
	for j := 0; j < len(universe[0]); j++ {
		isEmpty := true
		for i := 0; i < len(universe); i++ {
			if universe[i][j] == "#" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyGalaxyColumns = append(emptyGalaxyColumns, j)
		}
	}

	return emptyGalaxyRows, emptyGalaxyColumns
}
