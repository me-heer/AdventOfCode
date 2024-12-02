package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	validOne := 0
	validTwo := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		spaceSplit := strings.Split(line, " ")
		ruleSplit := spaceSplit[0]
		ruleMin, _ := strconv.Atoi(strings.Split(ruleSplit, "-")[0])
		ruleMax, _ := strconv.Atoi(strings.Split(ruleSplit, "-")[1])
		character := string(spaceSplit[1][0])
		inputStr := spaceSplit[2]

		partOneResult := applyRule(inputStr, character, ruleMin, ruleMax, 1)
		if partOneResult {
			validOne++
		}
		partTwoResult := applyRule(inputStr, character, ruleMin, ruleMax, 2)
		if partTwoResult {
			validTwo++
		}
	}
	println("PART 1:", validOne)
	println("PART 2:", validTwo)
}

func applyRule(s string, char string, min int, max int, part int) bool {
	if part == 1 {
		matchCount := 0
		for i := 0; i < len(s); i++ {
			if s[i] == char[0] {
				matchCount++
			}
		}

		if matchCount >= min && matchCount <= max {
			return true
		}
		return false
	} else {
		matchCount := 0
		if s[min-1] == char[0] {
			matchCount++
		}
		if s[max-1] == char[0] {
			matchCount++
		}
		if matchCount == 1 {
			return true
		}
		return false
	}
}
