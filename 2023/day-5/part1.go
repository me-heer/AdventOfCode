package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1() {
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
	seedRanges := strings.Split(strings.TrimSpace(strings.Split(seedsLine, ":")[1]), " ")
	var seeds []string
	for i := 0; i < len(seedRanges); i += 2 {
		lower, _ := strconv.Atoi(seedRanges[i])
		upper, _ := strconv.Atoi(seedRanges[i+1])
		upper--

		for j := lower; j <= lower+upper; j++ {
			strJ := strconv.Itoa(j)
			seeds = append(seeds, strJ)
		}
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

	minSeed := math.MaxInt64
	for _, seed := range seeds {
		seedNum, _ := strconv.Atoi(seed)
		for _, conversionRate := range conversionRates {
			for _, sts := range conversionRate {
				destStart, _ := strconv.Atoi(strings.Split(sts, " ")[0])
				sourceStart, _ := strconv.Atoi(strings.Split(sts, " ")[1])
				rangeNumber, _ := strconv.Atoi(strings.Split(sts, " ")[2])

				if seedNum >= sourceStart && seedNum <= sourceStart+rangeNumber-1 {
					seedNum = destStart + (seedNum - sourceStart)
					break
				}
			}
		}
		println(seedNum)
		if seedNum <= minSeed {
			minSeed = seedNum
		}
	}
	println(minSeed)

}
