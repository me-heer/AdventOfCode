package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var trails = 0

type Point struct {
	row int
	col int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]int

	for scanner.Scan() {
		row := scanner.Text()
		var rowInt []int
		for i := 0; i < len(row); i++ {
			var number int
			if row[i] == '.' {
				number = -1
			} else {
				number, _ = strconv.Atoi(string(rune(row[i])))
			}
			rowInt = append(rowInt, number)
		}
		matrix = append(matrix, rowInt)
	}

	sum := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				visited := make(map[Point]bool)
				traverse(Point{row, col}, matrix, visited)
				sum += len(visited)
			}
		}
	}
	println(sum)
	println(trails)
}


func traverse(trailhead Point, matrix [][]int, visited map[Point]bool) {
	currRow := trailhead.row
	currCol := trailhead.col
	if matrix[currRow][currCol] == 9 {
		visited[Point{currRow, currCol}] = true
		trails++
	}
	if currRow > 0 && matrix[currRow][currCol]+1 == matrix[currRow-1][currCol] {
		trailhead.row = currRow - 1
		trailhead.col = currCol
		traverse(trailhead, matrix, visited)
	}

	if currCol > 0 && matrix[currRow][currCol]+1 == matrix[currRow][currCol-1] {
		trailhead.row = currRow
		trailhead.col = currCol - 1
		traverse(trailhead, matrix, visited)
	}

	if currCol < len(matrix[0])-1 && matrix[currRow][currCol]+1 == matrix[currRow][currCol+1] {
		trailhead.row = currRow
		trailhead.col = currCol + 1
		traverse(trailhead, matrix, visited)
	}

	if currRow < len(matrix)-1 && matrix[currRow][currCol]+1 == matrix[currRow+1][currCol] {
		trailhead.row = currRow + 1
		trailhead.col = currCol
		traverse(trailhead, matrix, visited)
	}
}
