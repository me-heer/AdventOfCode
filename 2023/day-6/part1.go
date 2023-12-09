package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func part1() {
	input, err := os.Open("day-6/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)

	timeInput, _ := reader.ReadString('\n')

	r, _ := regexp.Compile("([0-9]+)")
	var times []int
	for _, timeStr := range r.FindAllString(timeInput, -1) {
		time, _ := strconv.Atoi(timeStr)
		times = append(times, time)
	}

	distanceInput, _ := reader.ReadString('\n')

	r, _ = regexp.Compile("([0-9]+)")
	var distances []int
	for _, distanceStr := range r.FindAllString(distanceInput, -1) {
		distance, _ := strconv.Atoi(distanceStr)
		distances = append(distances, distance)
	}

	var allWins []int
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		wins := 0
		for j := 1; j <= time; j++ {
			maxDistance := (time - j) * j
			if maxDistance > distance {
				wins++
			}
		}
		allWins = append(allWins, wins)
	}
	multiplicationResult := 1
	for _, w := range allWins {
		multiplicationResult = multiplicationResult * w
	}
	println(multiplicationResult)
}
