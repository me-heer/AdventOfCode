package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("day-12/input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	var springs []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = strings.TrimSpace(line)
		springs = append(springs, line)
	}

	for i, spring := range springs {
		println(i)
		leftPart := strings.Split(spring, " ")[0]
		rightPart := strings.Split(spring, " ")[1]

		expandedRightPart := ""
		expandedLeftPart := ""
		for x := 0; x < 5; x++ {
			expandedLeftPart += leftPart
			expandedLeftPart += "?"
			expandedRightPart += rightPart
			expandedRightPart += ","
		}
		expandedLeftPart = expandedLeftPart[:len(expandedLeftPart)-1]
		expandedRightPart = expandedRightPart[:len(expandedRightPart)-1]

		newSpring := expandedLeftPart + " " + expandedRightPart
		// replace ? with . and # all combinations
		replaceAndCheck(newSpring)
	}

	println(validCount)

}

var validCount = 0

func replaceAndCheck(spring string) {
	input := strings.Split(spring, " ")[0]
	numbers := strings.Split(spring, " ")[1]
	if strings.Contains(input, "?") {
		i1 := strings.Replace(input, "?", ".", 1)
		i2 := strings.Replace(input, "?", "#", 1)
		replaceAndCheck(i1 + " " + numbers)
		replaceAndCheck(i2 + " " + numbers)
	} else {
		isValid(spring)
	}
}

func isValid(spring string) bool {
	numbers := strings.Split(strings.Split(spring, " ")[1], ",")
	hashes := strings.Split(strings.Split(spring, " ")[0], ".")
	var newHashes []string
	for _, hash := range hashes {
		if hash != "" {
			newHashes = append(newHashes, hash)
		}
	}

	valid := true
	if len(newHashes) == len(numbers) {
		for i, newHash := range newHashes {
			n, _ := strconv.Atoi(strings.TrimSpace(numbers[i]))
			if strings.Count(newHash, "#") != n {
				valid = false
			}
		}
	} else {
		valid = false
	}
	if valid {
		validCount++
	}
	return valid
}
