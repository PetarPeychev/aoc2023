package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var digitWords = map[string]int{
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

func firstDigit(text string) int {
	for i, char := range text {
		if char >= '0' && char <= '9' {
			num, _ := strconv.Atoi(string(char))
			return num
		}
		for word, digit := range digitWords {
			if strings.Contains(text[:i+1], word) {
				return digit
			}
		}
	}
	return -1
}

func lastDigit(text string) int {
	for i := len(text) - 1; i >= 0; i-- {
		char := text[i]
		if char >= '0' && char <= '9' {
			num, _ := strconv.Atoi(string(char))
			return num
		}
		for word, digit := range digitWords {
			if strings.Contains(text[i:], word) {
				return digit
			}
		}
	}
	return -1
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var sum int
	line := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		text := scanner.Text()

		num1 := firstDigit(text)
		if num1 == -1 {
			panic("No digits found in line " + text)
		}

		num2 := lastDigit(text)
		if num2 == -1 {
			panic("No digits found in line " + text)
		}
		println(line, text, num1, " + ", num2, " = ", fmt.Sprint(num1)+fmt.Sprint(num2))

		num, _ := strconv.Atoi(fmt.Sprint(num1) + fmt.Sprint(num2))
		sum += num
	}

	println("Final sum is", sum)
}
