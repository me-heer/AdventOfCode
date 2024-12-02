package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var record []string
	valid := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if applyRules(record) {
				valid++
			}
			record = []string{}
			continue
		}

		records := strings.Split(line, " ")
		record = append(record, records...)
	}
	if applyRules(record) {
		valid++
	}
	println(valid)
}

func applyRules(recordStr []string) bool {
	record := make(map[string]string)
	for _, r := range recordStr {
		lhs := strings.Split(r, ":")[0]
		rhs := strings.Split(r, ":")[1]
		record[lhs] = rhs
	}

	if val, ok := record["byr"]; !ok {
		return false
	} else {
		if len(val) != 4 {
			return false
		}
		year, _ := strconv.Atoi(val)
		if year >= 1920 && year <= 2002 {

		} else {
			return false
		}
	}
	if val, ok := record["iyr"]; !ok {
		return false
	} else {
		if len(val) != 4 {
			return false
		}
		year, err := strconv.Atoi(val)
		if err != nil {
			return false
		}

		if year >= 2010 && year <= 2020 {

		} else {
			return false
		}
	}

	if val, ok := record["eyr"]; !ok {
		return false
	} else {
		if len(val) != 4 {
			return false
		}
		year, err := strconv.Atoi(val)
		if err != nil {
			return false
		}

		if year >= 2020 && year <= 2030 {

		} else {
			return false
		}
	}

	if height, ok := record["hgt"]; !ok {
		return false
	} else {
		if strings.Contains(height, "cm") {
			height = strings.ReplaceAll(height, "cm", "")
			hgtInNumber, err := strconv.Atoi(height)
			if err != nil {
				return false
			}

			if hgtInNumber >= 150 && hgtInNumber <= 193 {

			} else {
				return false
			}

		} else if strings.Contains(height, "in") {
			height = strings.ReplaceAll(height, "in", "")
			hgtInNumber, err := strconv.Atoi(height)
			if err != nil {
				return false
			}

			if hgtInNumber >= 59 && hgtInNumber <= 76 {

			} else {
				return false
			}

		} else {
			return false
		}
	}

	if val, ok := record["hcl"]; !ok {
		return false
	} else {
		if len(val) != 7 {
			return false
		}
		if val[0] != '#' {
			return false
		}

		for i := 1; i <= 6; i++ {
			char := rune(val[i])
			isDigit := unicode.IsDigit(char)
			isLowerHex := char >= 'a' && char <= 'f'
			if !isDigit && !isLowerHex {
				return false
			}
		}
	}

	if val, ok := record["ecl"]; !ok {
		return false
	} else {
		switch val {
		case "amb":
			break
		case "blu":
			break
		case "brn":
			break
		case "gry":
			break
		case "grn":
			break
		case "hzl":
			break
		case "oth":
			break
		default:
			return false
		}
	}

	if val, ok := record["pid"]; !ok {
		return false
	} else {
		if len(val) != 9 {
			return false
		}
		_, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
	}
	return true
}
