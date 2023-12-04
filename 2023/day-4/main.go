package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var matches = make(map[int]int)

func main() {
	input, err := os.Open("day-4/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)

	cardId := 1
	for {
		card, err := reader.ReadString('\n')
		if err != nil && len(card) == 0 {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		card = strings.TrimRight(card, "\n")

		part1(card, 0, cardId)
		cardId++
	}

	part2(matches)
}

func part1(card string, sum int, cardIndex int) int {
	winningNumbers := strings.Split(strings.TrimSpace(strings.Split(strings.Split(card, ":")[1], "|")[0]), " ")
	numbersWeHave := strings.Split(strings.TrimSpace(strings.Split(strings.Split(card, ":")[1], "|")[1]), " ")
	totalMatches := 0
	for _, numberWeHave := range numbersWeHave {
		if numberWeHave == "" {
			continue
		}
		for _, winningNumber := range winningNumbers {
			if winningNumber == "" {
				continue
			}
			nwh, _ := strconv.Atoi(numberWeHave)
			wn, _ := strconv.Atoi(winningNumber)
			if nwh == wn {
				totalMatches++
				break
			}
		}
	}
	matches[cardIndex] = totalMatches
	sum += int(math.Pow(float64(2), float64(totalMatches-1)))
	return sum
}

var findMatchesResult = make(map[int]int)

func part2(matches map[int]int) {
	sum := 0
	for cardNumber, _ := range matches {
		sum += findMatches(matches, cardNumber)
	}
	println(sum + len(matches))
	println("CACHE HIT: ", cacheHit)
	println("RESULT COMPUTED: ", loopExecuted)
}

var cacheHit = 0
var loopExecuted = 0

func findMatches(matches map[int]int, cardNumber int) int {
	result, ok := findMatchesResult[cardNumber]
	if ok {
		cacheHit++
		return result
	}
	loopExecuted++
	m := matches[cardNumber]
	sum := m
	for i := 0; i < m; i++ {
		sum += findMatches(matches, cardNumber+i+1)
	}
	findMatchesResult[cardNumber] = sum
	return sum
}
