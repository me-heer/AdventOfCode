package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	scanner := bufio.NewScanner(file)
	var rules []string
	var updates [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			// X := strings.Split(line, "|")[0]
			// Y := strings.Split(line, "|")[1]
			rules = append(rules, line)
		}

		if strings.Contains(line, ",") {
			update := strings.Split(line, ",")
			updates = append(updates, update)
		}
	}

	sum := 0
	for _, update := range updates {
		var currentRules []string
		for _, rule := range rules {
			X := strings.Split(rule, "|")[0]
			Y := strings.Split(rule, "|")[1]
			if slices.Contains(update, X) && slices.Contains(update, Y) {
				currentRules = append(currentRules, rule)
			}
		}

		doesAllRulesPass := doesAllRulesPass(update, currentRules)

		if doesAllRulesPass {
			middlePage, _ := strconv.Atoi(update[len(update)/2])
			sum += middlePage
		} else {
			println("INCORRECT: ")
			for _, v := range update {
				print(v, " ")
			}
			println()

			println("RELEVANT RULES: ")
			for _, v := range currentRules {
				println(v)
			}

			swapAndTryAgain(update, currentRules)
		}
	}
	println("PART 1:", sum)
	println("PART 2:", failedSum)
}

var failedSum = 0

func swapAndTryAgain(update []string, currentRules []string) {
	println("SWAPPING")
	copyUpdate := make([]string, len(update))
	copy(copyUpdate, update)

	hasNotPassed := true

	for hasNotPassed {
		for _, rule := range currentRules {
			X := strings.Split(rule, "|")[0]
			Y := strings.Split(rule, "|")[1]
			xIndex := slices.Index(update, X)
			yIndex := slices.Index(update, Y)
			if xIndex > yIndex {
				temp := copyUpdate[xIndex]
				copyUpdate[xIndex] = copyUpdate[yIndex]
				copyUpdate[yIndex] = temp
				if doesAllRulesPass(copyUpdate, currentRules) {
					println("PASSED AFTER SWAP")
					for _, v := range copyUpdate {
						print(v, " ")
					}
					println()
					middlePage, _ := strconv.Atoi(copyUpdate[len(copyUpdate)/2])
					failedSum += middlePage
					hasNotPassed = false
					break
				}
			}
		}
	}
}

func doesAllRulesPass(update []string, currentRules []string) bool {
	doesAllRulePass := true
	for _, rule := range currentRules {
		X := strings.Split(rule, "|")[0]
		Y := strings.Split(rule, "|")[1]
		if slices.Index(update, X) > slices.Index(update, Y) {
			doesAllRulePass = false
			break
		}
	}
	return doesAllRulePass
}
