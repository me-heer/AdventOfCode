package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Label struct {
	name        string
	focalLength int
}

var boxes = make(map[int][]Label)

func main() {
	input, _ := os.Open("day-15/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var sequence string

	for scanner.Scan() {
		sequence += scanner.Text()
	}

	var inputs []string

	inputs = strings.Split(sequence, ",")
	for _, input := range inputs {
		if strings.Contains(input, "-") {
			name := strings.Split(input, "-")[0]
			resultingBoxNumber := Hash(name)
			alreadyPresentLabels := boxes[resultingBoxNumber]
			if len(alreadyPresentLabels) > 0 {
				removalIndex := -1
				for i, l := range alreadyPresentLabels {
					if l.name == name {
						removalIndex = i
						break
					}
				}
				if removalIndex != -1 {
					alreadyPresentLabels = append(alreadyPresentLabels[:removalIndex], alreadyPresentLabels[removalIndex+1:]...)
				}
			}
			boxes[resultingBoxNumber] = alreadyPresentLabels
		} else if strings.Contains(input, "=") {
			name := strings.Split(input, "=")[0]
			focalLength, _ := strconv.Atoi(strings.Split(input, "=")[1])
			resultingBoxNumber := Hash(name)
			alreadyPresentLabels := boxes[resultingBoxNumber]
			if len(alreadyPresentLabels) > 0 {
				foundIndex := -1
				for i, l := range alreadyPresentLabels {
					if l.name == name {
						foundIndex = i
						l.focalLength = focalLength
						break
					}
				}
				newLabel := Label{
					name:        name,
					focalLength: focalLength,
				}

				if foundIndex == -1 {
					alreadyPresentLabels = append(alreadyPresentLabels, newLabel)
				} else {
					alreadyPresentLabels[foundIndex] = newLabel
				}
			} else {
				alreadyPresentLabels = append(alreadyPresentLabels, Label{
					name:        name,
					focalLength: focalLength,
				})
			}
			boxes[resultingBoxNumber] = alreadyPresentLabels
		}
	}

	sum := 0
	for boxNumber, box := range boxes {
		labels := box
		for slotNumber, l := range labels {
			result := (1 + boxNumber) * (slotNumber + 1) * (l.focalLength)
			sum += result
		}
	}
	println(sum)

}
