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

func findCalibration() int {
	fileBuf, error := os.ReadFile("input.txt")

	check(error)

	fileContent := string(fileBuf)

	lines := strings.Split(fileContent, "\n")

	total := 0

	for _, line := range lines {
		digitsArr := regexp.MustCompile("[0-9]+").FindAllString(line, -1)

		digits := strings.Split(strings.Join(digitsArr, ""), "")

		firstDigitStr := digits[0]
		lastDigitStr := digits[len(digits)-1]

		finalDigits, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstDigitStr, lastDigitStr))

		total += finalDigits
	}

	return total
}

func main() {
	fmt.Println(findCalibration())
}
