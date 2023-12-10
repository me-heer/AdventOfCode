package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("day-9/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)

	var dataset []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = strings.TrimRight(line, "\n")
		dataset = append(dataset, line)
	}

	sum := 0
	for _, line := range dataset {
		strValues := strings.Split(line, " ")
		var intValues []int
		for _, strVal := range strValues {
			intValue, _ := strconv.Atoi(strVal)
			intValues = append(intValues, intValue)
		}
		sum += predictBeginning(intValues)
	}
	println(sum)
}

func predict(values []int) int {
	var allDiffsMan [][]int
	allDiffsMan = append(allDiffsMan, values)

	for {
		diffs := make([]int, len(values)-1)
		for i := 0; i < len(values)-1; i++ {
			diffs[i] = values[i+1] - values[i]
		}
		allDiffsMan = append(allDiffsMan, diffs)
		values = diffs
		// check if all diffs are 0 otherwise do it again
		foundAllZeroValues := true
		for _, val := range diffs {
			if val != 0 {
				foundAllZeroValues = false
				break
			}
		}
		if foundAllZeroValues {
			break
		}
	}

	allDiffsMan[len(allDiffsMan)-1] = append(allDiffsMan[len(allDiffsMan)-1], 0)

	for i := len(allDiffsMan) - 1; i > 0; i-- {
		a := allDiffsMan[i][len(allDiffsMan[i])-1]
		b := allDiffsMan[i-1][len(allDiffsMan[i-1])-1]
		allDiffsMan[i-1] = append(allDiffsMan[i-1], a+b)
	}
	return allDiffsMan[0][len(allDiffsMan[0])-1]
}
func predictBeginning(values []int) int {
	var allDiffsMan [][]int
	allDiffsMan = append(allDiffsMan, values)

	for {
		diffs := make([]int, len(values)-1)
		for i := 0; i < len(values)-1; i++ {
			diffs[i] = values[i+1] - values[i]
		}
		allDiffsMan = append(allDiffsMan, diffs)
		values = diffs
		// check if all diffs are 0 otherwise do it again
		foundAllZeroValues := true
		for _, val := range diffs {
			if val != 0 {
				foundAllZeroValues = false
				break
			}
		}
		if foundAllZeroValues {
			break
		}
	}

	allDiffsMan[len(allDiffsMan)-1] = append([]int{0}, allDiffsMan[len(allDiffsMan)-1]...)

	for i := len(allDiffsMan) - 1; i > 0; i-- {
		a := allDiffsMan[i][0]
		b := allDiffsMan[i-1][0]
		allDiffsMan[i-1] = append([]int{b - a}, allDiffsMan[i-1]...)
	}
	return allDiffsMan[0][0]
}
