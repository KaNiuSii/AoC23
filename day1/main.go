package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		var firstDigit, lastDigit rune
		firstDigitFound := false

		for _, runeVal := range line {
			if unicode.IsDigit(runeVal) {
				firstDigit = runeVal
				firstDigitFound = true
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lastDigit = rune(line[i])
				break
			}
		}

		if firstDigitFound {
			combined := string(firstDigit) + string(lastDigit)
			value, err := strconv.Atoi(combined)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			sum += value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Println(sum)
}
