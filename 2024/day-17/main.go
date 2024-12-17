package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

var regA, regB, regC int
var program []int
var programStr string

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Register") {
			registerName := strings.Split(strings.Split(line, " ")[1], ":")[0]
			switch registerName {
			case "A":
				regA, _ = strconv.Atoi(strings.Split(line, " ")[2])
			case "B":
				regB, _ = strconv.Atoi(strings.Split(line, " ")[2])
			case "C":
				regC, _ = strconv.Atoi(strings.Split(line, " ")[2])
			}
		}

		if strings.Contains(line, "Program") {
			programStr = strings.Split(line, " ")[1]
			nums := strings.Split(strings.Split(line, " ")[1], ",")
			for _, n := range nums {
				num, _ := strconv.Atoi(n)
				program = append(program, num)
			}
		}
	}
	println(programStr)

	// originalRegA := regA
	originalRegB := regB
	originalRegC := regC

	println("START")
	i := 0
	for {
		regA = i
		regB = originalRegB
		regC = originalRegC
		output := execute()
		if output == programStr {
			println("RESULT: ", i, output)
			break
		} else {
			println(i, output)
			newI := adjustIterations(output, programStr, i)
			if newI != i {
				i = newI
				continue
			}
		}
		i++
	}
}

var modifier = 2

func adjustIterations(output string, desired string, currIter int) int {
	outputNums := strings.Split(output, ",")
	if len(outputNums) < 2 {
		return currIter
	}
	desiredNums := strings.Split(desired, ",")

	matchCount := 0
	for i := 0; i < modifier; i++ {
		if outputNums[len(outputNums)-1-i] == desiredNums[len(desiredNums)-1-i] {
			matchCount++
		}
	}
	if matchCount == modifier {
		println("Last", matchCount, "matched. Adjusting iterations to:", currIter*8*8)
		modifier += 2
		return currIter * 8 * 8
	}
	return currIter
}

func execute() string {
	ip := 0
	output := ""
	for {
		if ip >= len(program) {
			break
		}
		switch program[ip] {
		case 0:
			result := int(regA / int(math.Pow(2, float64(parseComboOperand(program[ip+1])))))
			regA = result
			ip += 2
		case 1:
			regB = regB ^ program[ip+1]
			ip += 2
		case 2:
			regB = parseComboOperand(program[ip+1]) % 8
			ip += 2
		case 3:
			if regA == 0 {
				ip += 2
			} else {
				ip = program[ip+1]
			}
		case 4:
			regB = regB ^ regC
			ip += 2
		case 5:
			result := parseComboOperand(program[ip+1]) % 8
			output += strconv.Itoa(result) + ","
			ip += 2
		case 6:
			result := int(regA / int(math.Pow(2, float64(parseComboOperand(program[ip+1])))))
			regB = result
			ip += 2
		case 7:
			result := int(regA / int(math.Pow(2, float64(parseComboOperand(program[ip+1])))))
			regC = result
			ip += 2
		}
	}
	output = strings.TrimSuffix(output, ",")
	return output
}

func parseComboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	}
	println("Invalid Operand")
	return operand
}
