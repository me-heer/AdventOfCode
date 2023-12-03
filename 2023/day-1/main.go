package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.Open("day-1/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)
	sum := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		number := numberFromCurrentLine2(line)
		println(number)
		sum += number
	}
	println(sum)
}

var lettersToDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func numberFromCurrentLine2(line string) int {
	validDigits := []string{
		"1", "one",
		"2", "two",
		"3", "three",
		"4", "four",
		"5", "five",
		"6", "six",
		"7", "seven",
		"8", "eight",
		"9", "nine",
	}

	min, max := len(line), -1
	var minVal, maxVal string
	for _, validDigit := range validDigits {
		// parse for digits and letters together
		indexOfDigit := strings.Index(line, validDigit)
		lastIndexOfDigit := strings.LastIndex(line, validDigit)

		if indexOfDigit != -1 && indexOfDigit <= min {
			min = indexOfDigit
			minVal = lettersToDigits[validDigit]
		}
		if lastIndexOfDigit != -1 && lastIndexOfDigit >= max {
			max = lastIndexOfDigit
			maxVal = lettersToDigits[validDigit]
		}

	}

	number, _ := strconv.Atoi(minVal + maxVal)
	return number
}

func numberFromCurrentLine(line string) int {
	var leftPtr, rightPtr string
	for _, char := range line {
		if !unicode.IsDigit(char) {
			continue
		}
		if leftPtr == "" {
			leftPtr = string(char)
		}
		rightPtr = string(char) // rightPtr is always updated to the last digit
	}
	number, _ := strconv.Atoi(leftPtr + rightPtr)
	return number
}
