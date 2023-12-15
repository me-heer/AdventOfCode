package main

import (
	"bufio"
	"os"
	"strings"
)

func part1() {
	input, _ := os.Open("day-15/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var sequence string

	for scanner.Scan() {
		sequence += scanner.Text()
	}

	var inputs []string

	sum := 0
	inputs = strings.Split(sequence, ",")
	for _, input := range inputs {
		// run Hash algorithm for input
		result := Hash(input)
		sum += result
		println(result)
	}
	println(sum)

}

func Hash(input string) int {
	result := 0
	for _, c := range input {
		result = result + int(c)
		result = result * 17
		result = result % 256
	}
	return result
}
