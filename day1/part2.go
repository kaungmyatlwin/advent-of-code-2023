package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extractNumbersFromNumerals(s string) []int {
	numerals := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	numeralRegExpStr := ""
	intSlice := []int{}

	for k, _ := range numerals {
		numeralRegExpStr += k + "|"
	}

	for i := 0; i < len(s); i++ {
		numeralsArr := regexp.MustCompile("^("+numeralRegExpStr+"[0-9])").FindAllString(s[i:], -1)

		if len(numeralsArr) > 0 {
			foundNumeral := numeralsArr[0]
			num, convErr := strconv.Atoi(foundNumeral)

			if convErr == nil {
				intSlice = append(intSlice, num)
			} else {
				intSlice = append(intSlice, numerals[foundNumeral])
			}
		}
	}

	return intSlice
}

func findCalibration(content string) int {
	lines := strings.Split(content, "\n")

	total := 0

	for _, line := range lines {
		intSlice := extractNumbersFromNumerals(line)

		firstDigit := intSlice[0]
		lastDigit := intSlice[len(intSlice)-1]

		finalDigits, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))

		total += finalDigits
	}

	return total
}

func main() {
	fileBuf, error := os.ReadFile("input.txt")
	fileContent := string(fileBuf)

	check(error)

	result := findCalibration(fileContent)

	fmt.Println(result)
}
