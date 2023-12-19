package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const SIZE = 500

type Point struct {
	i int
	j int
}

var visited = make(map[Point]bool)
var pointsInside = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var directions []string
	for scanner.Scan() {
		directions = append(directions, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	grid := make([][]string, SIZE)
	for i := 0; i < SIZE; i++ {
		grid[i] = make([]string, SIZE)
	}

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			grid[i][j] = "."
		}
	}
	row, col := 190, 90
	for i, input := range directions {
		println(i)
		splitInput := strings.Split(input, " ")
		dir := splitInput[0]
		n, _ := strconv.Atoi(splitInput[1])
		grid, row, col = move(dir, n, row, col, grid)
	}
	printGrid(grid)

	floodFill(150, 150, grid)
	println(pointsInside)
}

func floodFill(r, c int, grid [][]string) {
	moveInAllDirs(r, c, grid)

}

func moveInAllDirs(r, c int, grid [][]string) {
	goUp(r, c, grid)
	goDown(r, c, grid)
	goLeft(r, c, grid)
	goRight(r, c, grid)
}

func goUp(r, c int, grid [][]string) {
	r--
	if grid[r][c] == "." && !visited[Point{r, c}] {
		visited[Point{r, c}] = true
		pointsInside++
	} else {
		return
	}
	moveInAllDirs(r, c, grid)
}

func goDown(r, c int, grid [][]string) {
	r++
	if grid[r][c] == "." && !visited[Point{r, c}] {
		visited[Point{r, c}] = true
		pointsInside++
	} else {
		return
	}
	moveInAllDirs(r, c, grid)
}

func goLeft(r, c int, grid [][]string) {
	c--
	if grid[r][c] == "." && !visited[Point{r, c}] {
		visited[Point{r, c}] = true
		pointsInside++
	} else {
		return
	}
	moveInAllDirs(r, c, grid)
}

func goRight(r, c int, grid [][]string) {
	c++
	if grid[r][c] == "." && !visited[Point{r, c}] {
		visited[Point{r, c}] = true
		pointsInside++
	} else {
		return
	}
	moveInAllDirs(r, c, grid)
}

func printGrid(grid [][]string) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			print(grid[i][j])
		}
		println()
	}

	println("---")
}

func move(dir string, n int, r int, c int, grid [][]string) ([][]string, int, int) {
	switch dir {
	case "U":
		var i int
		for i = 1; i <= n; i++ {
			r--
			grid[r][c] = "#"
			pointsInside++
		}
	case "D":
		var i int
		for i = 1; i <= n; i++ {
			r++
			grid[r][c] = "#"
			pointsInside++
		}
	case "L":
		var j int
		for j = 1; j <= n; j++ {
			c--
			grid[r][c] = "#"
			pointsInside++
		}
	case "R":
		var j int
		for j = 1; j <= n; j++ {
			c++
			grid[r][c] = "#"
			pointsInside++
		}
	}
	return grid, r, c
}
