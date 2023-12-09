package main

import (
	"bufio"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var cardStrength = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

var orderByHandType = make(map[int][]string)

func part1() {
	input, err := os.Open("day-7/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)

	var handsStr []string
	for {
		hand, err := reader.ReadString('\n')
		if err != nil && len(hand) == 0 {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		hand = strings.TrimRight(hand, "\n")
		handsStr = append(handsStr, hand)
	}

	for _, handWithBid := range handsStr {
		// check which type it belongs to
		hand := strings.Split(handWithBid, " ")[0]

		unique := make(map[string]int)
		for i := 0; i < 5; i++ {
			unique[string(hand[i])] = unique[string(hand[i])] + 1
		}

		if len(unique) == 1 {
			orderByHandType[FIVE_OF_A_KIND] = append(orderByHandType[FIVE_OF_A_KIND], handWithBid)
		}
		if len(unique) == 2 {
			//four of a kind or full house

			// number of occurrences of each unique card in four of a kind: 4, 1
			// number of occurrences of each unique card in full house: 2, 3
			// differentiate from above by multiplying and checking

			multiplicationResult := 1
			for _, cardLabelCount := range unique {
				multiplicationResult = multiplicationResult * cardLabelCount
			}

			if multiplicationResult == 4 {
				orderByHandType[FOUR_OF_A_KIND] = append(orderByHandType[FOUR_OF_A_KIND], handWithBid)
			} else if multiplicationResult == 6 {
				orderByHandType[FULL_HOUSE] = append(orderByHandType[FULL_HOUSE], handWithBid)
			}
		}
		if len(unique) == 3 {
			//three of a kind or two pair

			// number of occurrences of each unique card in three of a kind: 3, 1, 1
			// number of occurrences of each unique card in two pair: 2, 2, 1
			// differentiate from above by multiplying and checking

			multiplicationResult := 1
			for _, cardLabelCount := range unique {
				multiplicationResult = multiplicationResult * cardLabelCount
			}

			if multiplicationResult == 3 {
				orderByHandType[THREE_OF_A_KIND] = append(orderByHandType[THREE_OF_A_KIND], handWithBid)
			} else if multiplicationResult == 4 {
				orderByHandType[TWO_PAIR] = append(orderByHandType[TWO_PAIR], handWithBid)
			}
		}
		if len(unique) == 4 {
			orderByHandType[ONE_PAIR] = append(orderByHandType[ONE_PAIR], handWithBid)
		}
		if len(unique) == 5 {
			orderByHandType[HIGH_CARD] = append(orderByHandType[HIGH_CARD], handWithBid)
		}
	}

	var orderedHandTypes []int
	for key, handsInHandType := range orderByHandType {
		orderedHandTypes = append(orderedHandTypes, key)
		handCmp := func(handA, handB string) int {
			// handA < handB return -1
			// handA > handB return +1
			for charIndex := 0; charIndex < 5; charIndex++ {
				strengthA := cardStrength[string(handA[charIndex])]
				strengthB := cardStrength[string(handB[charIndex])]
				if strengthA < strengthB {
					return -1
				} else if strengthA > strengthB {
					return +1
				}
			}
			return 0
		}
		slices.SortFunc(handsInHandType, handCmp)
	}
	sort.Ints(orderedHandTypes)

	sum := 0
	rank := 1
	for _, i := range orderedHandTypes {
		for _, hand := range orderByHandType[i] {
			bid, _ := strconv.Atoi(strings.Split(hand, " ")[1])
			sum = sum + (rank * bid)
			rank++
		}
	}
	println(sum)

}
