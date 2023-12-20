package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var workflows = make(map[string]Workflow)

type Workflow struct {
	rules []string
}

var parts = make([]Part, 0)

type Part struct {
	categories map[string]int
}

func main() {
	file, _ := os.Open("day-19/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		ruleStart := strings.Index(line, "{")
		ruleEnd := strings.Index(line, "}")
		workflowName := line[0:ruleStart]

		ruleStr := line[ruleStart+1 : ruleEnd]
		rules := strings.Split(ruleStr, ",")

		workflows[workflowName] = Workflow{rules: rules}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		ruleStart := strings.Index(line, "{")
		ruleEnd := strings.Index(line, "}")

		ruleStr := line[ruleStart+1 : ruleEnd]
		values := strings.Split(ruleStr, ",")

		var p Part
		p.categories = make(map[string]int)
		for _, v := range values {
			propertyName := strings.Split(v, "=")[0]
			propertyValue, _ := strconv.Atoi(strings.Split(v, "=")[1])
			p.categories[propertyName] = propertyValue
		}
		parts = append(parts, p)
	}

	for _, part := range parts {
		// start with in
		w := workflows["in"]
		executeWorkflow(w, part)
	}

	println(accepted)
	println(rejected)
	println(sum)
}

var accepted, rejected = 0, 0
var sum = 0

func executeWorkflow(workflow Workflow, part Part) {
	for _, r := range workflow.rules {
		if len(strings.Split(r, ":")) > 1 {
			// not terminal
			ruleCondition := strings.Split(r, ":")[0]
			destination := strings.Split(r, ":")[1]

			if strings.Contains(ruleCondition, "<") {
				propertyName := strings.Split(ruleCondition, "<")[0]
				propertyRequiredValue, _ := strconv.Atoi(strings.Split(ruleCondition, "<")[1])

				partValue := part.categories[propertyName]
				if partValue < propertyRequiredValue {
					// check if destination is workflow or terminal
					if len(destination) > 1 {
						// not terminal
						executeWorkflow(workflows[destination], part)
						break
					} else if len(destination) == 1 {
						// terminal
						if destination == "A" {
							accepted++

							for _, c := range part.categories {
								sum += c
							}

							break
						} else if destination == "R" {
							rejected++
							break
						} else {
							log.Fatal("NEITHER A OR R WAS FOUND")
						}
					} else {
						log.Fatal("INVALID DESTINATION LENGTH")
					}
				} else {
					continue
				}

			} else if strings.Contains(ruleCondition, ">") {
				propertyName := strings.Split(ruleCondition, ">")[0]
				propertyRequiredValue, _ := strconv.Atoi(strings.Split(ruleCondition, ">")[1])

				partValue := part.categories[propertyName]
				if partValue > propertyRequiredValue {
					// check if destination is workflow or terminal
					if len(destination) > 1 {
						// not terminal
						executeWorkflow(workflows[destination], part)
						break
					} else if len(destination) == 1 {
						// terminal
						if destination == "A" {
							accepted++
							for _, c := range part.categories {
								sum += c
							}
							break
						} else if destination == "R" {
							rejected++
							break
						} else {
							log.Fatal("NEITHER A OR R WAS FOUND")
						}
					} else {
						log.Fatal("INVALID DESTINATION LENGTH")
					}
				} else {
					continue
				}

			} else {
				log.Fatal("NO CONDITION TYPE MATCHED")
			}
		} else if len(r) == 1 {
			// terminal
			if r == "A" {
				accepted++
				for _, c := range part.categories {
					sum += c
				}
				break
			} else if r == "R" {
				rejected++
				break
			} else {
				log.Fatal("NEITHER A OR R WAS FOUND")
			}
		} else if len(r) > 1 {
			// non terminal
			executeWorkflow(workflows[r], part)
			break

		} else {
			log.Fatal("SHOULDN'T HAPPEN")
		}
	}
}
