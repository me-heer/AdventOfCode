package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var freeBlocks []Block

type Block struct {
	index int
	size  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var diskMap string
	for scanner.Scan() {
		diskMap = scanner.Text()
	}

	expandedDiskMap := expandDiskMap(diskMap)

	start := 0
	end := len(expandedDiskMap) - 1
	for start < end {
		if expandedDiskMap[start] == -1 && expandedDiskMap[end] != -1 {
			expandedDiskMap[start], expandedDiskMap[end] = expandedDiskMap[end], expandedDiskMap[start]
			start++
			end--
		} else if expandedDiskMap[start] != -1 {
			start++
		} else if expandedDiskMap[end] == -1 {
			end--
		}
	}

	checksum := int64(0)
	for i := 0; i < len(expandedDiskMap); i++ {
		if expandedDiskMap[i] == -1 {
			break
		}
		fileId := expandedDiskMap[i]
		checksum += (int64(i) * int64(fileId))
	}
	println(checksum)
}

func expandDiskMap(diskMap string) []int {
	var expandedDiskMap []int
	fileId := 0
	diskIndex := 0
	for i := 0; i < len(diskMap); i++ {
		if i%2 == 0 {
			fileSize, _ := strconv.Atoi(string(rune(diskMap[i])))
			for j := 0; j < fileSize; j++ {
				expandedDiskMap = append(expandedDiskMap, fileId)
				diskIndex++
			}
			fileId++
		} else {
			blockSize, _ := strconv.Atoi(string(rune(diskMap[i])))
			freeBlocks = append(freeBlocks, Block{index: diskIndex, size: blockSize})
			for j := 0; j < blockSize; j++ {
				expandedDiskMap = append(expandedDiskMap, -1)
				diskIndex++
			}
		}
	}
	return expandedDiskMap
}
