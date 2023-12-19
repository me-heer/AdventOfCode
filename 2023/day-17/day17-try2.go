package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

type AnotherPriorityQueue []*Node

func (pq AnotherPriorityQueue) Len() int { return len(pq) }

func (pq AnotherPriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss

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

	heatLoss int //used as priority in the priority queue

	// for finding out which direction we are headed
	directionRow           int
	directionCol           int
	nConsecutiveDirections int

	index int // used by heap
}

var grid [][]Node

type SeenInfo struct {
	row                    int
	col                    int
	directionRow           int
	directionCol           int
	nConsecutiveDirections int
}

type Direction struct {
	directionI int
	directionJ int
}

func main() {
	input, _ := os.Open("day-17/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var matrix []string

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	directions := make([]Direction, 0)
	directions = append(directions, Direction{0, 1})
	directions = append(directions, Direction{1, 0})
	directions = append(directions, Direction{0, -1})
	directions = append(directions, Direction{-1, 0})

	grid = make([][]Node, len(matrix))

	for i := 0; i < len(matrix); i++ {
		row := make([]Node, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			heatLoss, _ := strconv.Atoi(string(matrix[i][j]))
			row[j] = Node{
				row:      i,
				col:      j,
				heatLoss: heatLoss,
			}
		}
		grid[i] = row
	}

	pq := make(AnotherPriorityQueue, 0)
	source := Node{
		row:                    0,
		col:                    0,
		heatLoss:               0,
		directionRow:           0,
		directionCol:           0,
		nConsecutiveDirections: 0,
	}

	heap.Push(&pq, &source)
	seen := make(map[SeenInfo]bool)

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(*Node)

		if curr.row == len(grid)-1 && curr.col == len(grid[0])-1 && curr.nConsecutiveDirections >= 4 {
			println(curr.heatLoss)
			println("BREAKING")
			break
		}

		// Check if we have encountered this state before
		// Improvement over previous solution: we need to check how many consecutive steps you took
		currSeen := SeenInfo{
			row:                    curr.row,
			col:                    curr.col,
			directionRow:           curr.directionRow,
			directionCol:           curr.directionCol,
			nConsecutiveDirections: curr.nConsecutiveDirections,
		}

		if _, ok := seen[currSeen]; ok {
			continue
		}

		seen[currSeen] = true

		if curr.nConsecutiveDirections < 10 && !(curr.directionRow == 0 && curr.directionCol == 0) {
			rowForward := curr.row + curr.directionRow
			colForward := curr.col + curr.directionCol
			if rowForward < len(grid) && rowForward >= 0 && colForward >= 0 && colForward < len(grid[0]) {
				heap.Push(&pq, &Node{
					row:                    rowForward,
					col:                    colForward,
					heatLoss:               curr.heatLoss + grid[rowForward][colForward].heatLoss,
					directionRow:           curr.directionRow,
					directionCol:           curr.directionCol,
					nConsecutiveDirections: curr.nConsecutiveDirections + 1,
				})
			}
		}

		// go in other directions
		if curr.nConsecutiveDirections >= 4 || (curr.row == 0 && curr.col == 0) {
			for _, dir := range directions {
				if !(dir.directionI == curr.directionRow && dir.directionJ == curr.directionCol) && !(dir.directionI == -curr.directionRow && dir.directionJ == -curr.directionCol) {
					newDirI := curr.row + dir.directionI
					newDirJ := curr.col + dir.directionJ
					if newDirI < len(grid) && newDirI >= 0 && newDirJ >= 0 && newDirJ < len(grid[0]) {
						heap.Push(&pq, &Node{
							row:                    newDirI,
							col:                    newDirJ,
							heatLoss:               curr.heatLoss + grid[newDirI][newDirJ].heatLoss,
							directionRow:           dir.directionI,
							directionCol:           dir.directionJ,
							nConsecutiveDirections: 1,
						})
					}
				}
			}

		}

	}

}
