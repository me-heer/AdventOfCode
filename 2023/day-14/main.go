package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	input, _ := os.Open("day-14/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var matrix []string

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println("INPUT")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			print(string(matrix[i][j]))
		}
		println()
	}

	for i := 0; i < 1000; i++ {
		matrix = cycle(matrix)
		if i == 0 {
			sum := calculateSum(matrix)
			println(i, " ", sum)
			continue
		}

		sum := calculateSum(matrix)
		println(i, " ", sum)
	}

	sum := calculateSum(matrix)
	println(sum)
}

func calculateSum(matrix []string) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if string(matrix[i][j]) == "O" {
				sum += len(matrix) - i
			}
		}
	}
	return sum
}

func cycle(matrix []string) []string {
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if string(matrix[i][j]) == "O" {
				matrix = goNorthAndUpdate(i, j, matrix)
			}
		}
	}

	for j := 1; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			if string(matrix[i][j]) == "O" {
				matrix = goWestAndUpdate(i, j, matrix)
			}
		}
	}

	for i := len(matrix) - 2; i >= 0; i-- {
		for j := 0; j < len(matrix[0]); j++ {
			if string(matrix[i][j]) == "O" {
				matrix = goSouthAndUpdate(i, j, matrix)
			}
		}
	}

	for j := len(matrix[0]) - 2; j >= 0; j-- {
		for i := 0; i < len(matrix); i++ {
			if string(matrix[i][j]) == "O" {
				matrix = goEastAndUpdate(i, j, matrix)
			}
		}
	}
	return matrix
}

func goNorthAndUpdate(i int, j int, matrix []string) []string {
	farthestNorthI := goNorth(i, j, matrix)
	updatedCurrentRow := []rune(matrix[i])
	updatedCurrentRow[j] = '.'
	matrix[i] = string(updatedCurrentRow)

	updatedNorthRow := []rune(matrix[farthestNorthI])
	updatedNorthRow[j] = 'O'
	matrix[farthestNorthI] = string(updatedNorthRow)
	return matrix
}

func goWestAndUpdate(i int, j int, matrix []string) []string {
	farthestWestJ := goWest(i, j, matrix)
	updatedCurrentRow := []rune(matrix[i])
	updatedCurrentRow[j] = '.'
	updatedCurrentRow[farthestWestJ] = 'O'
	matrix[i] = string(updatedCurrentRow)
	return matrix
}

func goEastAndUpdate(i int, j int, matrix []string) []string {
	farthestEastJ := goEast(i, j, matrix)
	updatedCurrentRow := []rune(matrix[i])
	updatedCurrentRow[j] = '.'
	updatedCurrentRow[farthestEastJ] = 'O'
	matrix[i] = string(updatedCurrentRow)
	return matrix
}

func goSouthAndUpdate(i int, j int, matrix []string) []string {
	farthestSouthI := goSouth(i, j, matrix)
	updatedCurrentRow := []rune(matrix[i])
	updatedCurrentRow[j] = '.'
	matrix[i] = string(updatedCurrentRow)

	updatedSouthRow := []rune(matrix[farthestSouthI])
	updatedSouthRow[j] = 'O'
	matrix[farthestSouthI] = string(updatedSouthRow)
	return matrix
}

func goEast(currI, currJ int, matrix []string) int {
	var j int
	for j = currJ + 1; j < len(matrix[0]); j++ {
		if string(matrix[currI][j]) == "." {
			continue
		}
		return j - 1
	}
	return j - 1
}

func goWest(currI, currJ int, matrix []string) int {
	var j int
	for j = currJ - 1; j >= 0; j-- {
		if string(matrix[currI][j]) == "." {
			continue
		}
		return j + 1
	}
	return j + 1
}

func goSouth(currI, currJ int, matrix []string) int {
	var i int
	for i = currI + 1; i < len(matrix); i++ {
		if string(matrix[i][currJ]) == "." {
			continue
		}
		return i - 1
	}
	return i - 1
}

func goNorth(currI, currJ int, matrix []string) int {
	var i int
	for i = currI - 1; i >= 0; i-- {
		if string(matrix[i][currJ]) == "." {
			continue
		}
		return i + 1
	}
	return i + 1
}
