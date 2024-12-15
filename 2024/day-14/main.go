package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var width = 101
var height = 103

type Robot struct {
	xPos int
	yPos int
	xVel int
	yVel int
}

var fileCounter = 1 // Counter for output file names

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var robots []Robot
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`-?\d+`)
		matches := re.FindAllString(line, -1)

		xPos, _ := strconv.Atoi(matches[0])
		yPos, _ := strconv.Atoi(matches[1])

		xVel, _ := strconv.Atoi(matches[2])
		yVel, _ := strconv.Atoi(matches[3])
		robots = append(robots, Robot{xPos, yPos, xVel, yVel})
		println(xPos, yPos, xVel, yVel)
	}

	// reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(robots); j++ {
			robot := robots[j]
			robot.xPos = (robot.xPos + robot.xVel)
			if robot.xPos >= 0 {
				robot.xPos = robot.xPos % width
			} else {
				robot.xPos = width + robot.xPos
				robot.xPos = robot.xPos % width
			}

			robot.yPos = (robot.yPos + robot.yVel)
			if robot.yPos >= 0 {
				robot.yPos = robot.yPos % height
			} else {
				robot.yPos = height + robot.yPos
				robot.yPos = robot.yPos % height
			}
			robots[j] = robot
		}

		printRobots(robots)
		// _, _ = reader.ReadString('\n')
	}

	firstQuadrant := 0
	secondQuadrant := 0
	thirdQuadrant := 0
	fourthQuadrant := 0
	for r := 0; r < len(robots); r++ {
		robot := robots[r]
		if robot.yPos >= 0 && robot.yPos <= 50 && robot.xPos >= 0 && robot.xPos <= 49 {
			firstQuadrant++
		}
		if robot.yPos >= 0 && robot.yPos <= 50 && robot.xPos >= 51 && robot.xPos <= 100 {
			secondQuadrant++
		}
		if robot.yPos >= 52 && robot.yPos <= 102 && robot.xPos >= 0 && robot.xPos <= 49 {
			thirdQuadrant++
		}
		if robot.yPos >= 52 && robot.yPos <= 102 && robot.xPos >= 51 && robot.xPos <= 100 {
			fourthQuadrant++
		}
	}

	println(firstQuadrant, secondQuadrant, thirdQuadrant, fourthQuadrant)

	sum := firstQuadrant * secondQuadrant * thirdQuadrant * fourthQuadrant
	println(sum)
}

func printRobots(robots []Robot) {
	// Create a new file for each call
	fileName := fmt.Sprintf("%d.txt", fileCounter)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			found := false
			for _, r := range robots {
				if r.xPos == x && r.yPos == y {
					writer.WriteString("*")
					found = true
					break
				}
			}
			if !found {
				writer.WriteString(".")
			}
		}
		writer.WriteString("\n")
	}

	writer.Flush() // Ensure all data is written to the file
	fileCounter++  // Increment the file counter for the next call
}
