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

	regexpStr := ""
	intSlice := []int{}

	for k, _ := range numerals {
		regexpStr += k + "|"
	}

	numeralsArr := regexp.MustCompile(regexpStr+"([0-9])").FindAllString(s, -1)

	for _, k := range numeralsArr {
		num := numerals[strings.ToLower(k)]

		if num == 0 {
			kValArr := strings.Split(k, "")
			kVal := kValArr[len(kValArr)-1]
			i, _ := strconv.Atoi(kVal)

			intSlice = append(intSlice, i)
		} else {
			intSlice = append(intSlice, num)
		}
	}

	return intSlice
}

func findCalibration() int {
	fileBuf, error := os.ReadFile("input.txt")

	check(error)

	fileContent := string(fileBuf)

	lines := strings.Split(fileContent, "\n")

	total := 0

	for _, line := range lines {
		intSlice := extractNumbersFromNumerals(line)

		firstDigit := intSlice[0]
		lastDigit := intSlice[len(intSlice)-1]

		finalDigits, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		fmt.Println(line, intSlice, firstDigit, lastDigit, finalDigits, total)

		total += finalDigits
	}

	return total
}

func main() {
	result := findCalibration()

	fmt.Println(result)
}
