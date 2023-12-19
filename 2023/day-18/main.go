package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var directions []string
	for scanner.Scan() {
		directions = append(directions, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	const SIZE = 500

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
		println(input, i)
		splitInput := strings.Split(input, " ")
		dir := splitInput[0]
		n, _ := strconv.Atoi(splitInput[1])
		grid, row, col = move(dir, n, row, col, grid)
	}
	printGrid(grid)
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
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
		}
	case "D":
		var i int
		for i = 1; i <= n; i++ {
			r++
			grid[r][c] = "#"
		}
	case "L":
		var j int
		for j = 1; j <= n; j++ {
			c--
			grid[r][c] = "#"
		}
	case "R":
		var j int
		for j = 1; j <= n; j++ {
			c++
			grid[r][c] = "#"
		}
	}
	return grid, r, c
}
