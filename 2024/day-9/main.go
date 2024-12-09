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

	expandedDiskMap, maxFileId := expandDiskMap(diskMap)

	end := len(expandedDiskMap) - 1
	for maxFileId > 0 {
		fileSize := 0
		var currFileIndices []int

		for expandedDiskMap[end] != maxFileId {
			end--
		}

		for expandedDiskMap[end] == maxFileId {
			currFileIndices = append(currFileIndices, end)
			end--
			fileSize++
		}
		if fileSize == 0 {
			maxFileId--
			continue
		}

		moved := false
		for f := 0; f < len(freeBlocks); f++ {
			if freeBlocks[f].size >= fileSize && freeBlocks[f].index < end {
				moved = true
				for i := 0; i < fileSize; i++ {
					expandedDiskMap[freeBlocks[f].index] = maxFileId
					freeBlocks[f].index++
				}
				freeBlocks[f].size = freeBlocks[f].size - fileSize
				break
			}
		}

		if moved {
			for _, v := range currFileIndices {
				expandedDiskMap[v] = -1
			}
		}

		maxFileId--
	}

	fmt.Println(expandedDiskMap)

	checksum := int64(0)
	for i := 0; i < len(expandedDiskMap); i++ {
		if expandedDiskMap[i] == -1 {
			continue
		}
		fileId := expandedDiskMap[i]
		checksum += (int64(i) * int64(fileId))
	}
	println(checksum)
}

func expandDiskMap(diskMap string) ([]int, int) {
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
	return expandedDiskMap, fileId - 1
}
