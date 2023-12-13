package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("day-13/input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	patterns := make([][]string, 100)
	patternIndex := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				line = strings.TrimSpace(line)
				patterns[patternIndex] = append(patterns[patternIndex], line)
				break
			}
			panic(err)
		}
		line = strings.TrimSpace(line)
		if len(line) < 1 {
			patternIndex++
			continue
		}
		patterns[patternIndex] = append(patterns[patternIndex], line)
	}

	for i, pattern := range patterns {
		println("PROGRESS: ", i)
		reflect(pattern)
	}
	println(sum)
}

var sum = 0

func reflect(pattern []string) {
	result := reflectHorizontally(pattern)
	if result {
		return
	}
	reflectVertically(pattern)
}

func reflectHorizontally(pattern []string) bool {
	for i := 1; i < len(pattern); i++ {
		// we found a match
		result, smudgesUsed := tryHorizontalMatch(pattern, i-1, i)
		if result && smudgesUsed == 1 {
			println("FOUND MATCH AT ", i-1, " ", i)
			sum += ((i - 1) + 1) * 100
			return true
		}
	}
	return false
}

func tryHorizontalMatch(pattern []string, fromUp int, fromDown int) (bool, int) {
	isMatch := true
	smudgesUsed := 0
	for i, j := fromUp, fromDown; i >= 0 && j < len(pattern); i, j = i-1, j+1 {

		if pattern[i] != pattern[j] {
			matches := 0
			for c := 0; c < len(pattern[i]); c++ {
				if pattern[i][c] == pattern[j][c] {
					matches++
				}
			}

			if matches == len(pattern[i])-1 {
				smudgesUsed++
			} else {
				isMatch = false
				break
			}
		}

	}
	return isMatch, smudgesUsed
}

func tryVerticalMatch(pattern []string, fromLeft int, fromRight int) (bool, int) {
	isMatch := true
	smudgesUsed := 0
	for j1, j2 := fromLeft, fromRight; j1 >= 0 && j2 < len(pattern[0]); j1, j2 = j1-1, j2+1 {
		p1 := ""
		for i := 0; i < len(pattern); i++ {
			p1 += string(pattern[i][j1])
		}

		p2 := ""
		for i := 0; i < len(pattern); i++ {
			p2 += string(pattern[i][j2])
		}

		if p1 != p2 {
			matches := 0
			for c := 0; c < len(p1); c++ {
				if p1[c] == p2[c] {
					matches++
				}
			}

			if matches == len(p1)-1 {
				smudgesUsed++
			} else {
				isMatch = false
				break
			}
		}
	}
	return isMatch, smudgesUsed
}

func reflectVertically(pattern []string) {
	for j := 1; j < len(pattern[0]); j++ {
		p1 := ""
		for i := 0; i < len(pattern); i++ {
			p1 += string(pattern[i][j-1])
		}

		p2 := ""
		for i := 0; i < len(pattern); i++ {
			p2 += string(pattern[i][j])
		}

		result, smudgesUsed := tryVerticalMatch(pattern, j-1, j)
		if result && smudgesUsed == 1 {
			println("FOUND VERTICAL MATCH AT ", j-1, " ", j)
			sum += ((j - 1) + 1)
			break
		}
	}
}
