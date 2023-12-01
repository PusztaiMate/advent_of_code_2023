package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	cont := string(content)

	numbers := make([]string, 0)

	for i, line := range strings.Split(cont, "\n") {
		if line == "" {
			continue
		}
		lineBefore := line

		numberInLine := parseFirstAndLastNumberWordInString(line)

		numbers = append(numbers, numberInLine)
		fmt.Println(i, ":", lineBefore, "->", line, "->", numberInLine)
	}

	fmt.Println(sumUpDigits(numbers))
}

func parseFirstAndLastNumberWordInString(line string) string {
	line = replaceNumberWordsWithNumbers(line)

	firstDigit, err := findFirstDigit(line)
	if err != nil {
		fmt.Println("could not find shit in", line)
	}
	lastDigit, err := findLastDigit(line)
	if err != nil {
		fmt.Println("could not find shit in", line)
	}

	if firstDigit == "" {
		panic("something is not right")
	}

	return firstDigit + lastDigit
}

func replaceNumberWordsWithNumbers(input string) string {
	m := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	input = replaceFirstWord(input, m)
	input = replaceLastWord(input, m)

	return input
}

func replaceFirstWord(input string, m map[string]string) string {
	smallestKey := ""
	smallestValue := 999999
	digitsAsText := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for _, k := range digitsAsText {
		idx := strings.Index(input, k)
		if idx != -1 && idx < smallestValue {
			smallestKey = k
			smallestValue = idx
		}
	}

	if smallestKey != "" {
		input = strings.Replace(input, smallestKey, m[smallestKey], 1)
	}

	return input
}

func replaceLastWord(input string, m map[string]string) string {
	biggestKey := ""
	biggestValue := -1
	digitsAsText := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for _, k := range digitsAsText {
		idx := indexOfLastOccurrence(input, k)
		fmt.Println("looking at", k, "in", input, "and found idx", idx)
		if idx != -1 && idx > biggestValue {
			biggestKey = k
			biggestValue = idx
		}
	}

	if biggestKey != "" {
		ss := input[biggestValue:]
		fmt.Println("biggestKey is", biggestKey, "and is", ss)
		ssNew := strings.ReplaceAll(ss, biggestKey, m[biggestKey])
		input = strings.Replace(input, ss, ssNew, 1)
	}

	return input
}

func indexOfLastOccurrence(str, substr string) int {
	ss := str
	index := 0
	for {
		idx := strings.Index(ss, substr)
		if idx == -1 {
			break
		}

		index += idx
		ss = ss[idx+len(substr):]
	}

	return index
}

func sumUpDigits(input []string) int {
	numberOfOperands := 0
	sum := 0
	for _, number := range input {
		asInt, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		sum += asInt
		numberOfOperands++
	}

	fmt.Println("made", numberOfOperands, "additions")

	return sum
}

func findFirstDigit(line string) (string, error) {
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			return string(ch), nil
		}
	}

	return "", fmt.Errorf("could not find digit")
}

func findLastDigit(line string) (string, error) {
	lastDigit := ""
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			lastDigit = string(ch)
		}
	}

	if lastDigit != "" {
		return lastDigit, nil
	}

	return "", fmt.Errorf("could not find digit")
}
