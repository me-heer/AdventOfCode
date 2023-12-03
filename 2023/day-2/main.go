package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("day-2/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)

	const (
		RedCount   = 12
		GreenCount = 13
		BlueCount  = 14
	)
	sum := 0

	for {
		gameLine, err := reader.ReadString('\n')
		if err != nil && len(gameLine) == 0 {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		//gameId, _ := strconv.Atoi(strings.Split(strings.Split(gameLine, ":")[0], " ")[1])
		gameSets := strings.Split(strings.TrimSpace(strings.Split(gameLine, ":")[1]), ";")
		maxCubeCount := make(map[string]int)
		for _, gameSet := range gameSets {
			cubeCount := make(map[string]int)
			cubes := strings.Split(gameSet, ",")
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				count, _ := strconv.Atoi(strings.Split(cube, " ")[0])
				color := strings.Split(cube, " ")[1]
				cubeCount[color] = cubeCount[color] + count
			}

			for color, count := range cubeCount {
				if count > maxCubeCount[color] {
					maxCubeCount[color] = count
				}
			}
		}
		power := 1
		for _, count := range maxCubeCount {
			power *= count
		}

		sum += power
	}
	println(sum)
}
