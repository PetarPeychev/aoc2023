package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	regexGame       = regexp.MustCompile(`Game (\d+): `)
	regexReveal     = regexp.MustCompile(`((?:\d+) (?:red|green|blue)(?:, (?:\d+) (?:red|green|blue))*)`)
	regexRevealItem = regexp.MustCompile(`(\d+) (red|green|blue)`)
)

type Reveal struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Number  int
	Reveals []Reveal
}

func parseGame(str string) (game Game) {
	gameNumber, _ := strconv.Atoi(regexGame.FindStringSubmatch(str)[1])
	game.Number = gameNumber

	reveals := regexReveal.FindAllStringSubmatch(str, -1)
	for _, reveal := range reveals {
		revealStr := reveal[1]
		revealItems := regexRevealItem.FindAllStringSubmatch(revealStr, -1)
		var reveal Reveal
		for _, revealItem := range revealItems {
			number, _ := strconv.Atoi(revealItem[1])
			switch revealItem[2] {
			case "red":
				reveal.Red = number
			case "green":
				reveal.Green = number
			case "blue":
				reveal.Blue = number
			}
		}
		game.Reveals = append(game.Reveals, reveal)
	}

	return game
}

func (game Game) isValid(red int, green int, blue int) bool {
	for _, reveal := range game.Reveals {
		if reveal.Red > red || reveal.Green > green || reveal.Blue > blue {
			return false
		}
	}
	return true
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
		game := parseGame(text)
		if game.isValid(12, 13, 14) {
			sum += game.Number
		}
		fmt.Printf("%+v - %v\n\n", game, game.isValid(12, 13, 14))
	}

	println("Final sum is", sum)
}
