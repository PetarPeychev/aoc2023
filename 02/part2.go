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

func (game Game) minCubes() (red int, green int, blue int) {
	minRed := 0
	minGreen := 0
	minBlue := 0
	for _, reveal := range game.Reveals {
		if reveal.Red > minRed {
			minRed = reveal.Red
		}
		if reveal.Green > minGreen {
			minGreen = reveal.Green
		}
		if reveal.Blue > minBlue {
			minBlue = reveal.Blue
		}
	}
	return minRed, minGreen, minBlue
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
		red, green, blue := game.minCubes()
		power := red * green * blue
		sum += power
		fmt.Printf("Game %d: %d cubes\n", game.Number, power)
	}

	println("Final sum is", sum)
}
