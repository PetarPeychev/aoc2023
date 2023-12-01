package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var sum int
	line := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		text := scanner.Text()
		var digits []rune
		for _, char := range text {
			if char >= '0' && char <= '9' {
				digits = append(digits, char)
			}
		}
		print(line, " ", text, " - ", string(digits), " - ")
		if len(digits) == 0 {
			panic("No digits found in line " + text)
		}
		digit1 := digits[0]
		digit2 := digits[len(digits)-1]

		num, err := strconv.Atoi(string(digit1) + string(digit2))
		if err != nil {
			panic(err)
		}
		sum += num
		println(fmt.Sprint(num), " = ", sum)
	}

	println("Final sum is", sum)
}
