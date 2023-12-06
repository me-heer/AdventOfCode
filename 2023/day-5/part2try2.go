package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("day-5/input.txt")
	if err != nil {
		panic(err)
	}

	defer func(input *os.File) {
		err := input.Close()
		if err != nil {
			panic(err)
		}
	}(input)

	reader := bufio.NewReader(input)
	seedsLine, err := reader.ReadString('\n')
	seedRangesLine := strings.Split(strings.TrimSpace(strings.Split(seedsLine, ":")[1]), " ")
	var seedRanges [][]int
	for i := 0; i < len(seedRangesLine); i += 2 {
		lower, _ := strconv.Atoi(seedRangesLine[i])
		upper, _ := strconv.Atoi(seedRangesLine[i+1])
		upper--
		rangeInt := []int{lower, lower + upper}
		seedRanges = append(seedRanges, rangeInt)
	}

	var inputLines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = strings.TrimRight(line, "\n")
		inputLines = append(inputLines, line)
	}

	var seedToSoil []string
	var soilToFertilizer []string
	var fertilizerToWater []string
	var waterToLight []string
	var lightToTemperature []string
	var temperatureToHumidity []string
	var humidityToLocation []string

	for lineIndex, line := range inputLines {
		if strings.Contains(line, "seed-to-soil") {
			for i := lineIndex + 1; inputLines[i] != ""; i++ {
				seedToSoil = append(seedToSoil, inputLines[i])
			}
		}
		if strings.Contains(line, "soil-to-fertilizer") {
			for i := lineIndex + 1; inputLines[i] != ""; i++ {
				soilToFertilizer = append(soilToFertilizer, inputLines[i])
			}
		}
		if strings.Contains(line, "fertilizer-to-water") {
			for i := lineIndex + 1; inputLines[i] != ""; i++ {
				fertilizerToWater = append(fertilizerToWater, inputLines[i])
			}
		}
		if strings.Contains(line, "water-to-light") {
			for i := lineIndex + 1; inputLines[i] != ""; i++ {
				waterToLight = append(waterToLight, inputLines[i])
			}
		}
		if strings.Contains(line, "light-to-temperature") {
			for i := lineIndex + 1; inputLines[i] != ""; i++ {
				lightToTemperature = append(lightToTemperature, inputLines[i])
			}
		}
		if strings.Contains(line, "temperature-to-humidity") {
			for i := lineIndex + 1; inputLines[i] != ""; i++ {
				temperatureToHumidity = append(temperatureToHumidity, inputLines[i])
			}
		}
		if strings.Contains(line, "humidity-to-location") {
			for i := lineIndex + 1; i < len(inputLines); i++ {
				humidityToLocation = append(humidityToLocation, inputLines[i])
			}
		}
	}

	var conversionRates [][]string
	conversionRates = append(conversionRates, seedToSoil)
	conversionRates = append(conversionRates, soilToFertilizer)
	conversionRates = append(conversionRates, fertilizerToWater)
	conversionRates = append(conversionRates, waterToLight)
	conversionRates = append(conversionRates, lightToTemperature)
	conversionRates = append(conversionRates, temperatureToHumidity)
	conversionRates = append(conversionRates, humidityToLocation)

	//answer := math.MaxInt64
	for _, seedRange := range seedRanges {
		var currentRanges [][]int
		currentRanges = append(currentRanges, seedRange)
		for _, conversionRate := range conversionRates {
			var executedRanges [][]int
			for _, sts := range conversionRate {
				destLower, _ := strconv.Atoi(strings.Split(sts, " ")[0])
				sourceLower, _ := strconv.Atoi(strings.Split(sts, " ")[1])
				rangeNumber, _ := strconv.Atoi(strings.Split(sts, " ")[2])

				// max between sourceLower and currentLower
				// min between sourceUpper and currentUpper

				var resultRanges [][]int
				for _, currentRange := range currentRanges {
					currentLower := currentRange[0]
					currentUpper := currentRange[1]
					sourceUpper := sourceLower + rangeNumber - 1

					if (currentLower > sourceUpper && currentUpper > sourceUpper) || (currentUpper < sourceLower && currentLower < sourceLower) {
						resultRanges = append(resultRanges, currentRange)
						continue
					}

					alreadyExecuted := false
					for _, executedRange := range executedRanges {
						if currentLower == executedRange[0] && currentUpper == executedRange[1] {
							alreadyExecuted = true
							break
						}
					}
					if alreadyExecuted {
						resultRanges = append(resultRanges, currentRange)
						break
					}

					var lowerMax int
					if currentLower >= sourceLower {
						lowerMax = currentLower
					} else {
						lowerMax = sourceLower
					}

					var upperMin int
					if currentUpper <= sourceUpper {
						upperMin = currentUpper
					} else {
						upperMin = sourceUpper
					}

					//[lowerMax, upperMin]
					//println("currentLower: ", currentLower)
					//println("currentHigher: ", currentUpper)
					//println("lowerMax:", lowerMax)
					//println("upperMin:", upperMin)

					if currentLower < sourceLower {
						var newRange = []int{currentLower, sourceLower - 1}
						resultRanges = append(resultRanges, newRange)
					}
					if currentUpper > sourceUpper {
						var newRange = []int{sourceUpper + 1, currentUpper}
						resultRanges = append(resultRanges, newRange)
					}

					lowerMax = destLower + (lowerMax - sourceLower)
					upperMin = destLower + (upperMin - sourceLower)

					var newRange = []int{lowerMax, upperMin}
					resultRanges = append(resultRanges, newRange)
					executedRanges = append(executedRanges, newRange)
				}
				currentRanges = resultRanges
			}
		}
		for _, cr := range currentRanges {
			println(cr[0], " ", cr[1])
		}
		// find minimum number here
	}
}
