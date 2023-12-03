package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PartNumber struct {
	Number int
	Line   int
	Start  int
	End    int
}

type Symbol struct {
	Line int
	Pos  int
}

func parseParts(text string, line int) []PartNumber {
	var numbers []PartNumber
	digits := ""
	for i, c := range text {
		if c >= '0' && c <= '9' {
			digits += string(c)
		} else if digits != "" {
			num, _ := strconv.Atoi(digits)
			partNum := PartNumber{num, line, i - len(digits), i - 1}
			numbers = append(numbers, partNum)
			digits = ""
		}
	}
	if digits != "" {
		num, _ := strconv.Atoi(digits)
		partNum := PartNumber{num, line, len(text) - len(digits), len(text) - 1}
		numbers = append(numbers, partNum)
	}
	return numbers
}

func parseSymbols(text string, line int) []Symbol {
	var symbols []Symbol
	for i, c := range text {
		if c != '.' && (c < '0' || c > '9') {
			symbols = append(symbols, Symbol{line, i})
		}
	}
	return symbols
}

func main() {
	file, _ := os.Open("input1.txt")
	defer file.Close()

	var sum int
	line := 0
	var lineParts []PartNumber
	var lineSymbols [][]Symbol

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parts := parseParts(text, line)
		fmt.Printf("%+v", parts)
		lineParts = append(lineParts, parts...)
		symbols := parseSymbols(text, line)
		fmt.Printf("%+v\n", symbols)
		lineSymbols = append(lineSymbols, symbols)
		line++
	}

	for _, part := range lineParts {
		if part.Line != 0 {
			for _, symbol := range lineSymbols[part.Line-1] {
				if symbol.Pos >= part.Start-1 && symbol.Pos <= part.End+1 {
					sum += part.Number
					break
				}
			}
		}
		for _, symbol := range lineSymbols[part.Line] {
			if symbol.Pos == part.Start-1 || symbol.Pos == part.End+1 {
				sum += part.Number
				break
			}
		}
		if part.Line < len(lineSymbols)-1 {
			for _, symbol := range lineSymbols[part.Line+1] {
				if symbol.Pos >= part.Start-1 && symbol.Pos <= part.End+1 {
					sum += part.Number
					break
				}
			}
		}
	}

	println("Final sum is", sum)
}
