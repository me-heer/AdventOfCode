package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
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

	//move := network["AAA"]

	// find all nodes which end with A
	var startNodes []string
	var moves [][]string
	for node, nextNodes := range network {
		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
			moves = append(moves, nextNodes)
		}
	}

	// for each node, find path to reach end node
	// since this path has periodicity,
	//we can find the least common multiple of the result for each node and that will be the answer
	var lcm []int
	for i := range startNodes {
		directionIndex := 0
		steps := 0
		for {
			dir := string(directions[directionIndex])
			steps++

			var nextNode string
			if dir == "L" {
				nextNode = moves[i][0]
			} else {
				nextNode = moves[i][1]
			}

			moves[i] = network[nextNode]
			if strings.HasSuffix(nextNode, "Z") {
				break
			}

			directionIndex++
			if directionIndex == len(directions) {
				directionIndex = 0
			}
		}
		lcm = append(lcm, steps)
	}

	// find lcm among all numbers in lcm array
	result := 0
	for i := 0; i < len(lcm)-1; i++ {
		result = findLcm(lcm[i], lcm[i+1])
		lcm[i+1] = result
	}
	println(result)

}

func findLcm(a, b int) int {
	return (a * b) / findGcd(a, b)
}

func findGcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
