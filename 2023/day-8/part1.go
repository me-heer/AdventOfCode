package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func part1() {
	input, err := os.Open("day-8/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)
	directions, _ := reader.ReadString('\n')
	directions = strings.TrimRight(directions, "\n")

	_, _ = reader.ReadString('\n')

	network := make(map[string][]string)

	for {
		nodeLine, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		node := strings.Split(nodeLine, " = ")[0]
		dirs := strings.Split(nodeLine, " = ")[1]
		left := dirs[1:4]
		right := dirs[6:9]
		network[node] = []string{left, right}
	}

	move := network["AAA"]
	directionIndex := 0
	steps := 0
	for {
		dir := string(directions[directionIndex])
		steps++

		var nextNode string
		if dir == "L" {
			nextNode = move[0]
		} else {
			nextNode = move[1]
		}
		move = network[nextNode]

		if nextNode == "ZZZ" {
			break
		}

		directionIndex++
		if directionIndex == len(directions) {
			directionIndex = 0
		}
	}
	println(steps)

}
