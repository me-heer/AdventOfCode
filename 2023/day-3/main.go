package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbers []string

func main() {
	input, err := os.Open("day-3/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)

	// Reading into matrix
	var matrix [][]string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = strings.TrimRight(line, "\n")

		row := make([]string, len(line))
		for i, char := range line {
			row[i] = string(char)
		}
		matrix = append(matrix, row)
	}

	sum := 0

	for rowIndex, row := range matrix {
		for colIndex, col := range row {
			if !unicode.IsDigit(rune(col[0])) && col != "." {
				var adjacentNumbers []string
				if rowIndex > 0 {
					topRowIndex := rowIndex - 1

					topLeftColIndex := colIndex - 1
					leftPtr := goLeft(topLeftColIndex, matrix, topRowIndex)
					rightPtr := goRight(topLeftColIndex, matrix, topRowIndex)
					adjacentNumber := extractAdjacentNumber(leftPtr, rightPtr, matrix, topRowIndex)
					if adjacentNumber != "" {
						adjacentNumbers = append(adjacentNumbers, adjacentNumber)
					}

					topColIndex := colIndex
					if rightPtr < topColIndex {
						leftPtr = goLeft(topColIndex, matrix, topRowIndex)
						rightPtr = goRight(topColIndex, matrix, topRowIndex)
						adjacentNumber = extractAdjacentNumber(leftPtr, rightPtr, matrix, topRowIndex)
						if adjacentNumber != "" {
							adjacentNumbers = append(adjacentNumbers, adjacentNumber)
						}
					}

					topRightColIndex := colIndex + 1
					if rightPtr < topRightColIndex {
						leftPtr = goLeft(topRightColIndex, matrix, topRowIndex)
						rightPtr = goRight(topRightColIndex, matrix, topRowIndex)
						adjacentNumber = extractAdjacentNumber(leftPtr, rightPtr, matrix, topRowIndex)
						if adjacentNumber != "" {
							adjacentNumbers = append(adjacentNumbers, adjacentNumber)
						}
					}
				}
				currentRowIndex := rowIndex

				leftColIndex := colIndex - 1
				leftPtr := goLeft(leftColIndex, matrix, currentRowIndex)
				rightPtr := goRight(leftColIndex, matrix, currentRowIndex)
				adjacentNumber := extractAdjacentNumber(leftPtr, rightPtr, matrix, currentRowIndex)
				if adjacentNumber != "" {
					adjacentNumbers = append(adjacentNumbers, adjacentNumber)
				}

				rightColIndex := colIndex + 1
				leftPtr = goLeft(rightColIndex, matrix, currentRowIndex)
				rightPtr = goRight(rightColIndex, matrix, currentRowIndex)
				adjacentNumber = extractAdjacentNumber(leftPtr, rightPtr, matrix, currentRowIndex)
				if adjacentNumber != "" {
					adjacentNumbers = append(adjacentNumbers, adjacentNumber)
				}

				if rowIndex < len(matrix)-1 {
					bottomRowIndex := rowIndex + 1

					bottomLeftColIndex := colIndex - 1
					leftPtr = goLeft(bottomLeftColIndex, matrix, bottomRowIndex)
					rightPtr = goRight(bottomLeftColIndex, matrix, bottomRowIndex)
					adjacentNumber = extractAdjacentNumber(leftPtr, rightPtr, matrix, bottomRowIndex)
					if adjacentNumber != "" {
						adjacentNumbers = append(adjacentNumbers, adjacentNumber)
					}

					bottomColIndex := colIndex
					if rightPtr < bottomColIndex {
						leftPtr = goLeft(bottomColIndex, matrix, bottomRowIndex)
						rightPtr = goRight(bottomColIndex, matrix, bottomRowIndex)
						adjacentNumber = extractAdjacentNumber(leftPtr, rightPtr, matrix, bottomRowIndex)
						if adjacentNumber != "" {
							adjacentNumbers = append(adjacentNumbers, adjacentNumber)
						}
					}

					bottomRightColIndex := colIndex + 1
					if rightPtr < bottomRightColIndex {
						leftPtr = goLeft(bottomRightColIndex, matrix, bottomRowIndex)
						rightPtr = goRight(bottomRightColIndex, matrix, bottomRowIndex)
						adjacentNumber = extractAdjacentNumber(leftPtr, rightPtr, matrix, bottomRowIndex)
						if adjacentNumber != "" {
							adjacentNumbers = append(adjacentNumbers, adjacentNumber)
						}
					}
				}
				if len(adjacentNumbers) == 2 {
					n1, _ := strconv.Atoi(adjacentNumbers[0])
					n2, _ := strconv.Atoi(adjacentNumbers[1])
					gearRatio := n1 * n2
					sum += gearRatio
				}
			}
		}
	}
	println(sum)
}

func findAdjacentNumberInDirection(colIndex int, matrix [][]string, rowIndex int) (int, int) {
	leftPtr := goLeft(colIndex, matrix, rowIndex)
	rightPtr := goRight(colIndex, matrix, rowIndex)
	extractAdjacentNumber(leftPtr, rightPtr, matrix, rowIndex)
	return leftPtr, rightPtr
}

func extractAdjacentNumber(leftPtr int, rightPtr int, matrix [][]string, rowIndex int) string {
	number := ""
	for i := leftPtr; i <= rightPtr; i++ {
		number += matrix[rowIndex][i]
	}
	if len(number) == 0 {
		return ""
	}
	return number
}

func goRight(topLeftColIndex int, matrix [][]string, topRowIndex int) int {
	var rightPtr int
	// go right
	for rightPtr = topLeftColIndex; rightPtr <= len(matrix[topRowIndex])-1 && unicode.IsDigit(rune(matrix[topRowIndex][rightPtr][0])); rightPtr++ {
	}
	rightPtr-- // to account for the extra right in the last iteration when condition fails
	return rightPtr
}

func goLeft(from int, matrix [][]string, topRowIndex int) int {
	// go left
	var leftPtr int
	for leftPtr = from; leftPtr >= 0 && unicode.IsDigit(rune(matrix[topRowIndex][leftPtr][0])); leftPtr-- {
	}
	leftPtr++ // to account for the extra left in the last iteration when condition fails
	return leftPtr
}
