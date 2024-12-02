package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into two parts
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to integer:", line)
			continue
		}
		left = append(left, num1)
		right = append(right, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(left)
	sort.Ints(right)

	distance := 0
	for i := 0; i < len(left); i++ {
		d := left[i] - right[i]
		if d < 0 {
			d = d * (-1)
		}
		distance += d
	}
	println("PART 1 :", distance)

	countMap := make(map[int]int)
	for _, num := range right {
		countMap[num]++
	}

	similarityScore := 0
	for i := 0; i < len(left); i++ {
		v, present := countMap[left[i]]
		if present {
			similarityScore += left[i] * v
		}
	}
	println(similarityScore)
}
