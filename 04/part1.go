package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	regexCard = regexp.MustCompile(`Card[ ]+(\d+):([^|]+)\|([^|]+)`)
)

type Card struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
}

func (card Card) calculatePoints() int {
	points := 0
	for _, winningNum := range card.WinningNumbers {
		for _, num := range card.Numbers {
			if winningNum == num {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
	}
	return points
}

func parseCard(str string) (card Card) {
	matches := regexCard.FindStringSubmatch(str)
	cardId, _ := strconv.Atoi(matches[1])
	card.Id = cardId

	winningNumStrings := strings.Split(matches[2], " ")
	for _, winningNumString := range winningNumStrings {
		if winningNumString != "" {
			winningNum, _ := strconv.Atoi(winningNumString)
			card.WinningNumbers = append(card.WinningNumbers, winningNum)
		}
	}

	numStrings := strings.Split(matches[3], " ")
	for _, numString := range numStrings {
		if numString != "" {
			num, _ := strconv.Atoi(numString)
			card.Numbers = append(card.Numbers, num)
		}
	}

	return card
}

func main() {
	file, _ := os.Open("input1.txt")
	defer file.Close()

	var sum int
	line := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		text := scanner.Text()
		card := parseCard(text)
		fmt.Printf("%+v", card)
		points := card.calculatePoints()
		fmt.Printf(" points: %d\n", points)
		sum += points
	}

	println("Final sum is", sum)
}
