package main

import (
	"bufio"
	"container/heap"
	"os"
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
	cost int
	// for finding out location
	row int
	col int
	// for finding out which direction we are headed
	directionRow int
	directionCol int
	index        int // used by heap
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

var matrix []string

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	var sr, sc int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 'S' {
				sr = i
				sc = j
				break
			}
		}
	}

	pq := make(AnotherPriorityQueue, 0)
	source := Node{
		row:          sr,
		col:          sc,
		cost:         0,
		directionRow: 0,
		directionCol: 1,
	}

	heap.Push(&pq, &source)
	seen := make(map[SeenInfo]bool)
	seen[SeenInfo{
		row: sr, col: sc, directionRow: 0, directionCol: 1,
	}] = true

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(*Node)
		// Check if we have encountered this state before
		currSeen := SeenInfo{
			row:          curr.row,
			col:          curr.col,
			directionRow: curr.directionRow,
			directionCol: curr.directionCol,
		}
		seen[currSeen] = true
		println("CURR: ", curr.row, curr.col, string(matrix[curr.row][curr.col]))
		if matrix[curr.row][curr.col] == 'E' {
			println("COST: ", curr.cost)
			break
		}

		// go forward
		// go counterclockwise
		// go clockwise

		dr := curr.directionRow
		dc := curr.directionCol
		if isInBounds(curr.row+dr, curr.col+dc) && matrix[curr.row+dr][curr.col+dc] != '#' {
			if seen[SeenInfo{row: curr.row + dr,
				col:          curr.col + dc,
				directionRow: dr,
				directionCol: dc,
			}] {
				continue
			}

			heap.Push(&pq, &Node{
				row:          curr.row + dr,
				col:          curr.col + dc,
				directionRow: dr,
				directionCol: dc,
				cost:         curr.cost + 1,
			})
		}
		dr = curr.directionCol
		dc = -curr.directionRow
		if matrix[curr.row][curr.col] != '#' {

			if seen[SeenInfo{row: curr.row,
				col:          curr.col,
				directionRow: dr,
				directionCol: dc,
			}] {
				continue
			}
			heap.Push(&pq, &Node{
				row:          curr.row,
				col:          curr.col,
				directionRow: dr,
				directionCol: dc,
				cost:         curr.cost + 1000,
			})
		}

		dr = -curr.directionCol
		dc = curr.directionRow
		if matrix[curr.row][curr.col] != '#' {
			if seen[SeenInfo{row: curr.row + dr,
				col:          curr.col + dc,
				directionRow: dr,
				directionCol: dc,
			}] {
				continue
			}
			heap.Push(&pq, &Node{
				row:          curr.row,
				col:          curr.col,
				directionRow: dr,
				directionCol: dc,
				cost:         curr.cost + 1000,
			})
		}

	}

}

func isInBounds(currRow int, currCol int) bool {
	return currRow >= 0 && currRow < len(matrix) && currCol >= 0 && currCol < len(matrix[0])
}
