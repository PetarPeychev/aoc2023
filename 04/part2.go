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

func (card Card) countMatches() int {
	matches := 0
	for _, winningNum := range card.WinningNumbers {
		for _, num := range card.Numbers {
			if winningNum == num {
				matches += 1
			}
		}
	}
	return matches
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

	var cardMatches []int
	line := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		text := scanner.Text()
		card := parseCard(text)
		fmt.Printf("%+v", card)
		matches := card.countMatches()
		fmt.Printf(" points: %d\n", matches)
		cardMatches = append(cardMatches, matches)
	}

	cardCopies := make([]int, len(cardMatches))
	for i := range cardCopies {
		cardCopies[i] = 1
	}
	for i, matches := range cardMatches {
		for j := 0; j < cardCopies[i]; j++ {
			for k := i + 1; k <= min(i+matches, len(cardMatches)-1); k++ {
				cardCopies[k] += 1
			}
		}
	}
	fmt.Printf("%+v\n", cardCopies)

	sum := 0
	for _, copies := range cardCopies {
		sum += copies
	}
	println("Final sum is", sum)
}
