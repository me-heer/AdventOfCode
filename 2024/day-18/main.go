package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

type AnotherPriorityQueue []*Node

func (pq AnotherPriorityQueue) Len() int { return len(pq) }

func (pq AnotherPriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost

}

func (pq AnotherPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *AnotherPriorityQueue) Push(x any) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *AnotherPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

type Node struct {
	// for finding out location
	row int
	col int

	cost int //used as priority in the priority queue

	// for finding out which direction we are headed
	directionRow int
	directionCol int

	index int // used by heap
}

type SeenInfo struct {
	row          int
	col          int
	directionRow int
	directionCol int
}

type Direction struct {
	directionI int
	directionJ int
}

var matrix [][]int
var matrixSize = 71
var bytes []string
var simulateBytes = 1024

var START = 1
var END = 2
var WALL = -1

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		bytes = append(bytes, line)
	}

	for i := 0; i < matrixSize; i++ {
		var row []int
		for j := 0; j < matrixSize; j++ {
			row = append(row, 0)
		}
		matrix = append(matrix, row)
	}

	directions := make([]Direction, 0)
	directions = append(directions, Direction{0, 1})
	directions = append(directions, Direction{1, 0})
	directions = append(directions, Direction{0, -1})
	directions = append(directions, Direction{-1, 0})

	for i := 0; i < len(bytes); i++ {
		b := bytes[i]
		x, _ := strconv.Atoi(strings.Split(b, ",")[0])
		y, _ := strconv.Atoi(strings.Split(b, ",")[1])
		matrix[y][x] = WALL
		//check

		pq := make(AnotherPriorityQueue, 0)
		source := Node{
			row:          0,
			col:          0,
			cost:         0,
			directionRow: 0,
			directionCol: 0,
		}

		heap.Push(&pq, &source)
		seen := make(map[SeenInfo]bool)
		pathExists := false

		for len(pq) > 0 {
			curr := heap.Pop(&pq).(*Node)

			// println("CURR: ", curr.row, curr.col)
			if curr.row == matrixSize-1 && curr.col == matrixSize-1 {
				pathExists = true
				// println(curr.cost)
				println("PATH EXISTS. ", i)
				break
			}

			// Check if we have encountered this state before
			// Improvement over previous solution: we need to check how many consecutive steps you took
			currSeen := SeenInfo{
				row:          curr.row,
				col:          curr.col,
				directionRow: curr.directionRow,
				directionCol: curr.directionCol,
			}

			if _, ok := seen[currSeen]; ok {
				// println("SEEN")
				continue
			}

			seen[currSeen] = true

			// go in other directions
			for _, dir := range directions {
				newDirI := curr.row + dir.directionI
				newDirJ := curr.col + dir.directionJ
				if newDirI < matrixSize && newDirI >= 0 && newDirJ >= 0 && newDirJ < matrixSize && matrix[newDirI][newDirJ] != WALL {
					// println("PUSHING: ", newDirI, newDirJ)
					heap.Push(&pq, &Node{
						row:          newDirI,
						col:          newDirJ,
						cost:         curr.cost + 1,
						directionRow: dir.directionI,
						directionCol: dir.directionJ,
					})
				}
			}
		}

		if !pathExists {
			println("PATH DOES NOT EXIST. ", b)
			break

		}
	}

}
