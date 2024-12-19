package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

var towels = make(map[string]bool)
var designs []string
var maxlen = 0

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			towelsStr := strings.Split(line, ", ")
			for _, t := range towelsStr {
				towels[t] = true
			}
		} else if line == "" {
			continue
		} else {
			designs = append(designs, line)
		}
	}

	for t, _ := range towels {
		if len(t) > maxlen {
			maxlen = len(t)
		}
	}

	sum := 0
	for _, d := range designs {
		sum += matchDesign(d)
	}

	println(sum)
}

var cache = make(map[string]int)

func matchDesign(remaining string) int {
	if len(remaining) == 0 {
		return 1
	}
	if v, ok := cache[remaining]; ok {
		return v
	}

	count := 0
	limit := int(math.Min(float64(len(remaining)), float64(maxlen)) + 1)
	for i := 0; i < limit; i++ {
		if towels[remaining[:i]] {
			c := matchDesign(remaining[i:])
			count += c
		}

	}
	cache[remaining] = count
	return count
}
