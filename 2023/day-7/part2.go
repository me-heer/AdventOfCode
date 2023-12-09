package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var updatedCardStrength = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

func main() {
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

		highestHandType := math.MinInt
		for card, _ := range updatedCardStrength {
			if card == "J" {
				continue
			}

			replacedHand := strings.ReplaceAll(hand, "J", card)
			unique := make(map[string]int)
			for i := 0; i < 5; i++ {
				unique[string(replacedHand[i])] = unique[string(replacedHand[i])] + 1
			}
			handType := determineHandType(unique)
			if handType > highestHandType {
				highestHandType = handType
			}
		}

		orderByHandType[highestHandType] = append(orderByHandType[highestHandType], handWithBid)
	}

	var orderedHandTypes []int
	for key, handsInHandType := range orderByHandType {
		orderedHandTypes = append(orderedHandTypes, key)
		handCmp := func(handA, handB string) int {
			// handA < handB return -1
			// handA > handB return +1
			for charIndex := 0; charIndex < 5; charIndex++ {
				strengthA := updatedCardStrength[string(handA[charIndex])]
				strengthB := updatedCardStrength[string(handB[charIndex])]
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

func determineHandType(unique map[string]int) int {
	if len(unique) == 1 {
		return FIVE_OF_A_KIND
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
			return FOUR_OF_A_KIND
		} else if multiplicationResult == 6 {
			return FULL_HOUSE
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
			return THREE_OF_A_KIND
		} else if multiplicationResult == 4 {
			return TWO_PAIR
		}
	}
	if len(unique) == 4 {
		return ONE_PAIR
	}
	if len(unique) == 5 {
		return HIGH_CARD
	}
	return -1
}
