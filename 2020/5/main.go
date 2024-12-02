package main

import (
	"bufio"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	maxSeat := -1
	seats := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		seat := findSeat(line)
		if seat > maxSeat {
			maxSeat = seat
		}
		seats = append(seats, seat)
	}
	sort.Ints(seats)
	for i, seat := range seats {
		var diff = 0
		if i != 0 {
			diff = seat - seats[i-1]
		}
		if diff != 1 {
			print("DIFF NOT ONE: ", seat)
		}
		println(seat, diff)
	}
}

func findSeat(input string) int {
	rowTop := 0
	rowBottom := 127
	for i := 0; i < 7; i++ {
		if input[i] == 'F' {
			middle := (rowTop + rowBottom) / 2
			rowBottom = middle
		} else if input[i] == 'B' {
			middle := (rowTop + rowBottom + 1) / 2
			rowTop = middle
		}
	}
	colLeft := 0
	colRight := 7
	for j := 7; j < 10; j++ {
		if input[j] == 'L' {
			middle := (colLeft + colRight) / 2
			colRight = middle
		} else if input[j] == 'R' {
			middle := (colLeft + colRight + 1) / 2
			colLeft = middle
		}
	}
	return (rowTop * 8) + colLeft
}
