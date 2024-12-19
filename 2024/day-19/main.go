package main

import (
	"bufio"
	"os"
	"strings"
)

var towels []string
var designs []string
var possible = make(map[string]int)
var cache = make(map[string]string)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			towels = strings.Split(line, ", ")
		} else if line == "" {
			continue
		} else {
			designs = append(designs, line)
		}
	}

	for _, t := range towels {
		println(t)
	}
	for _, d := range designs {
		println(d)
		matchDesign("", d, d)
	}
	sum := 0
	for _, v := range possible {
		sum += v
	}

	println(sum)
}

var noSolutionCache = make(map[string]bool)

func matchDesign(curr string, remaining string, final string) {
	if noSolutionCache[remaining] {
		println(remaining, " NOT POSSIBLE. RETURNING")
		return
	}
	if len(remaining) == 0 {
		println("FINAL: ", final)
		possible[final] += 1
		return
	}
	// println("Curr: ", curr, "Remaining: ", remaining, "FINAL: ", final)
	for _, t := range towels {
		if strings.HasPrefix(remaining, t) {
			newRemaining := remaining[0+len(t):]
			newCurr := curr + t
			if len(newRemaining) == 0 || newCurr == final {
				possible[final] += 1
				cache[remaining] = "END"
				break
			}
			if noSolutionCache[newRemaining] {
				continue
			}
			cache[remaining] = newRemaining
			matchDesign(newCurr, newRemaining, final)
		}
	}
	if possible[final] == 0 {
		println("NO SOLUTION FOUND. ", remaining)
		noSolutionCache[remaining] = true
	}
}
