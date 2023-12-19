package main

import (
	"bufio"
	"log"
	"math"
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
	file, _ := os.Open("day-18/input.txt")
	scanner := bufio.NewScanner(file)

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
	var points = make([]Point, 0)
	row, col := 0, 0
	boundaries := 0
	for _, input := range directions {
		splitInput := strings.Split(input, " ")
		hex := splitInput[2]
		hexNum := hex[2 : len(hex)-2]

		number, _ := strconv.ParseInt(hexNum, 16, 64)

		actualDir, _ := strconv.Atoi(hex[len(hex)-2 : len(hex)-1])
		var dir string
		switch actualDir {
		case 0:
			dir = "R"
		case 1:
			dir = "D"
		case 2:
			dir = "L"
		case 3:
			dir = "U"
		}
		n, _ := strconv.Atoi(splitInput[1])
		n = int(number)
		boundaries += n
		row, col = move(dir, n, row, col)
		points = append(points, Point{row, col})
	}

	sum1, sum2 := 0, 0
	for i := 0; i < len(points)-1; i++ {
		sum1 += points[i].i * points[i+1].j
		sum2 += points[i].j * points[i+1].i
	}
	result := int(math.Abs(float64(sum1-sum2))) / 2
	println(result)
	println(boundaries)
	println(result + boundaries/2 + 1)

}

func move(dir string, n int, r int, c int) (int, int) {
	switch dir {
	case "U":
		var i int
		for i = 1; i <= n; i++ {
			r--
			pointsInside++
		}
	case "D":
		var i int
		for i = 1; i <= n; i++ {
			r++
			pointsInside++
		}
	case "L":
		var j int
		for j = 1; j <= n; j++ {
			c--
			pointsInside++
		}
	case "R":
		var j int
		for j = 1; j <= n; j++ {
			c++
			pointsInside++
		}
	}
	return r, c
}
