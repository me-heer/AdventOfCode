package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	n     int
	times int
}

var blinkMap = make(map[Point]int)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stones []int

	for scanner.Scan() {
		row := scanner.Text()
		nums := strings.Split(row, " ")
		var rowInt []int
		for i := 0; i < len(nums); i++ {
			number, _ := strconv.Atoi(nums[i])
			rowInt = append(rowInt, number)
		}
		stones = rowInt
	}

	sum := 0
	for _, v := range stones {
		sum += blink(v, 75)
	}
	fmt.Println("SUM: ", sum)
}

func blink(n int, times int) int {
	digitsStr := strconv.Itoa(n)
	if v, ok := blinkMap[Point{n, times}]; ok {
		return v
	}
	if times == 1 {
		if n == 0 {
			return 1
		} else if len(digitsStr)%2 == 0 {
			return 2
		} else {
			return 1
		}
	} else {
		result := 0
		if n == 0 {
			result = blink(1, times-1)
		} else if len(digitsStr)%2 == 0 {
			left := digitsStr[0 : len(digitsStr)/2]
			right := digitsStr[len(digitsStr)/2:]

			leftInt, _ := strconv.Atoi(left)
			rightInt, _ := strconv.Atoi(right)
			result = blink(leftInt, times-1) + blink(rightInt, times-1)
		} else {
			result = blink(n*2024, times-1)
		}
		blinkMap[Point{n, times}] = result
		return result
	}
}
