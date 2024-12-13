package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	scanner := bufio.NewScanner(file)

	type Machine struct {
		Ax float64
		Ay float64
		Bx float64
		By float64
		Px float64
		Py float64
	}

	var machines []Machine
	i := 0
	var ax, ay, bx, by, px, py float64
	for scanner.Scan() {
		if i == 3 {
			scanner.Text()
			m := Machine{
				ax, ay, bx, by, px, py,
			}
			machines = append(machines, m)
			i = 0
			continue
		}
		line := scanner.Text()
		if strings.Contains(line, "Button A") {
			l1 := strings.Split(line, " ")[2]
			l2 := strings.Split(line, " ")[3]
			ax, _ = strconv.ParseFloat(strings.Split(strings.Split(l1, ",")[0], "+")[1], 32)
			ay, _ = strconv.ParseFloat(strings.Split(l2, "+")[1], 32)

		}
		if strings.Contains(line, "Button B") {
			l1 := strings.Split(line, " ")[2]
			l2 := strings.Split(line, " ")[3]
			bx, _ = strconv.ParseFloat(strings.Split(strings.Split(l1, ",")[0], "+")[1], 32)
			by, _ = strconv.ParseFloat(strings.Split(l2, "+")[1], 32)
		}
		if strings.Contains(line, "Prize") {
			px, _ = strconv.ParseFloat(strings.Split(strings.Split(strings.Split(line, " ")[1], ",")[0], "=")[1], 32)
			py, _ = strconv.ParseFloat(strings.Split(strings.Split(line, " ")[2], "=")[1], 32)
		}
		i++
	}
	m := Machine{
		ax, ay, bx, by, px, py,
	}
	machines = append(machines, m)

	sum := 0
	for i := 0; i < len(machines); i++ {
		m := machines[i]
		m.Px += 10000000000000
		m.Py += 10000000000000
		b := math.Round((((m.Px * m.Ay) / m.Ax) - m.Py) / (((m.Ay * m.Bx) / m.Ax) - m.By))
		a := math.Round((m.Py - (b * m.By)) / m.Ay)

		xResult := (a * m.Ax) + (b * m.Bx)
		yResult := (a * m.Ay) + (b * m.By)
		fmt.Println(i, a, b, "xResult: ", xResult, m.Px, "yResult: ", yResult, m.Py)
		if a == math.Trunc(a) {
			println("A TRUNC EQUAL")
		} else {
			continue
		}
		if b == math.Trunc(b) {
			println("B TRUNC EQUAL")
		} else {
			continue
		}
		if xResult == m.Px && yResult == m.Py {
			println(i, a, b, "PASS")
			sum += int((a * 3) + (b * 1))
		}
	}

	println(sum)

}

func roundToDecimal(num float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(num*factor) / factor
}
