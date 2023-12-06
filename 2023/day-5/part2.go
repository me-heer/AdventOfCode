package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func part2try1() {
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
	inputSeedRanges := strings.Split(strings.TrimSpace(strings.Split(seedsLine, ":")[1]), " ")
	var seedRanges [][2]string
	var seeds []string

	for i := 0; i < len(inputSeedRanges); i += 2 {
		lowerStr := inputSeedRanges[i]
		lower, _ := strconv.Atoi(inputSeedRanges[i])
		upper, _ := strconv.Atoi(inputSeedRanges[i+1])
		upper--
		upperStr := strconv.Itoa(lower + upper)
		seedRange := [2]string{lowerStr, upperStr}
		seedRanges = append(seedRanges, seedRange)
	}

	println("TOTAL SEEDS: ", len(seeds))

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

	for _, seedRange := range seedRanges {
		var currentRanges [][2]string
		currentRanges = append(currentRanges, seedRange)
		for _, conversionRate := range conversionRates {
			for _, sts := range conversionRate {
				destLower, _ := strconv.Atoi(strings.Split(sts, " ")[0])
				sourceLower, _ := strconv.Atoi(strings.Split(sts, " ")[1])
				rangeNumber, _ := strconv.Atoi(strings.Split(sts, " ")[2])

				var resultRanges [][2]string
				for _, currentRange := range currentRanges {
					currentLower, _ := strconv.Atoi(currentRange[0])
					currentUpper, _ := strconv.Atoi(currentRange[1])
					sourceUpper := sourceLower + rangeNumber - 1

					if currentLower >= sourceLower && currentUpper > sourceUpper && currentLower <= sourceUpper {
						// [currentLower, sourceUpper] with execution, [sourceUpper + 1, currentUpper]
						executedCurrentLower := strconv.Itoa(destLower + (currentLower - sourceLower))
						executedSourceUpper := strconv.Itoa(destLower + (sourceUpper - sourceLower))
						resultRange1 := [2]string{executedCurrentLower, executedSourceUpper}
						a2 := sourceUpper + 1
						b2 := currentUpper

						for _, rates := range conversionRate {
							destStart, _ := strconv.Atoi(strings.Split(rates, " ")[0])
							sourceStart, _ := strconv.Atoi(strings.Split(rates, " ")[1])
							rangeNumber, _ := strconv.Atoi(strings.Split(rates, " ")[2])

							if a2 >= sourceStart && a2 <= sourceStart+rangeNumber-1 {
								a2 = destStart + (a2 - sourceStart)
								if b2 >= sourceStart && b2 <= sourceStart+rangeNumber-1 {
									b2 = destStart + (b2 - sourceStart)
								}
								break
							}

						}

						resultRange2 := [2]string{strconv.Itoa(a2), strconv.Itoa(b2)}
						resultRanges = append(resultRanges, resultRange1, resultRange2)
					} else if currentLower < sourceLower && currentUpper >= sourceUpper {
						// [currentLower, sourceLower - 1], [sourceLower, currentUpper] with execution
						a1 := currentLower
						b1 := sourceLower - 1

						for _, rates := range conversionRate {
							destStart, _ := strconv.Atoi(strings.Split(rates, " ")[0])
							sourceStart, _ := strconv.Atoi(strings.Split(rates, " ")[1])
							rangeNumber, _ := strconv.Atoi(strings.Split(rates, " ")[2])

							if a1 >= sourceStart && a1 <= sourceStart+rangeNumber-1 {
								a1 = destStart + (a1 - sourceStart)
								if b1 >= sourceStart && b1 <= sourceStart+rangeNumber-1 {
									b1 = destStart + (b1 - sourceStart)
								}
								break
							}
						}

						resultRange1 := [2]string{strconv.Itoa(a1), strconv.Itoa(b1)}

						executedSourceLower := strconv.Itoa(destLower + (sourceLower - sourceLower))
						executedCurrentUpper := strconv.Itoa(destLower + (currentUpper - sourceLower))
						resultRange2 := [2]string{executedSourceLower, executedCurrentUpper}
						resultRanges = append(resultRanges, resultRange1, resultRange2)
					} else if currentLower >= sourceLower && currentUpper <= sourceUpper {
						executedCurrentLower := strconv.Itoa(destLower + (currentLower - sourceLower))
						executedCurrentUpper := strconv.Itoa(destLower + (currentUpper - sourceLower))
						resultRange := [2]string{executedCurrentLower, executedCurrentUpper}
						resultRanges = append(resultRanges, resultRange)
					} else {
						resultRange := currentRange
						resultRanges = append(resultRanges, resultRange)
					}
				}
				currentRanges = resultRanges
			}
		}
		println("TEST")
	}

}
