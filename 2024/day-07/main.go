package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var equations []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		equations = append(equations, line)
	}

	for i, equation := range equations {
		println(i)
		desiredResult, _ := strconv.Atoi(strings.Split(equation, ":")[0])
		numbers := strings.Split(strings.Split(equation, ": ")[1], " ")
		currResult, _ := strconv.Atoi(numbers[0])
		tryNumber(currResult, 0, numbers, desiredResult, 1)
		tryNumber(currResult, 1, numbers, desiredResult, 1)
		tryNumber(currResult, 2, numbers, desiredResult, 1)
	}
	sum := 0
	for i, _ := range count {
		sum += i
	}
	println(sum)
}

var count = make(map[int]bool)

func tryNumber(currResult int, operator int, numbers []string, desiredResult int, next int) {
	nextNum, _ := strconv.Atoi(numbers[next])
	if operator == 0 {
		currResult += nextNum
	} else if operator == 1 {
		currResult *= nextNum
	} else {
		result := strconv.Itoa(currResult)
		result = result + numbers[next]
		currResult, _ = strconv.Atoi(result)
	}
	next++
	if next >= len(numbers) {
		if currResult == desiredResult {
			count[desiredResult] = true
		}
		return
	}

	tryNumber(currResult, 0, numbers, desiredResult, next)
	tryNumber(currResult, 1, numbers, desiredResult, next)
	tryNumber(currResult, 2, numbers, desiredResult, next)
}
