package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.Open("day-6/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)

	timeInput, _ := reader.ReadString('\n')

	r, _ := regexp.Compile("([0-9]+)")
	timeNumberStr := ""
	for _, timeStr := range r.FindAllString(timeInput, -1) {
		timeNumberStr += timeStr
	}
	timeNumber, _ := strconv.Atoi(timeNumberStr)

	distanceInput, _ := reader.ReadString('\n')

	distanceNumberStr := ""
	for _, distanceStr := range r.FindAllString(distanceInput, -1) {
		distanceNumberStr += distanceStr
	}
	distanceNumber, _ := strconv.Atoi(distanceNumberStr)

	wins := 0
	for j := 1; j <= timeNumber; j++ {
		maxDistance := (timeNumber - j) * j
		if maxDistance > distanceNumber {
			wins++
		}
	}
	println(wins)
}
