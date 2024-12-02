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

	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		reports := make([]int, len(parts))
		for i, v := range parts {
			reports[i], err = strconv.Atoi(v)
			if err != nil {
				panic("Couldn't convert to int")
			}
		}


		if checkRules(reports) {
			counter++
		} else {
			for i := 0; i < len(reports); i++ {
				copySlice := make([]int, len(reports))
				copy(copySlice, reports)
				reportAfterOneRemoval := append(copySlice[:i], copySlice[i+1:]...)
				if checkRules(reportAfterOneRemoval) {
					counter++
					break
				} else {
					continue
				}
			}
		}
	}

	println(counter)
}

func checkRules(report []int) bool {
	decreasing := false
	increasing := false

	if report[0] > report[1] {
		decreasing = true
	} else if report[1] > report[0] {
		increasing = true
	} else if report[0] == report[1] {
		return false
	}

	for i := 1; i < len(report); i++ {
		if increasing {
			result := report[i-1] < report[i]
			if !result {
				return false
			}
		} else if decreasing {
			result := report[i-1] > report[i]
			if !result {
				return false
			}
		} else {
			return false
		}
	}

	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
