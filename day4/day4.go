package day4

import (
	"bufio"
	"github.com/fwielstra/AoC2023/utils"
	"io"
	"strings"
)

type Card struct {
	WinningNumbers []int
	OwnedNumbers   []int
}

type Cards []Card

func ParseCard(input string) Card {
	// discard Card x: prefix using indexing / slicing
	values := input[strings.Index(input, ":")+1:]
	winning, owned, _ := strings.Cut(values, "|")

	return Card{
		WinningNumbers: utils.TrimmedIntFields(winning),
		OwnedNumbers:   utils.TrimmedIntFields(owned),
	}
}

func ParseCards(r io.Reader) []Card {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	cards := make([]Card, 0)
	for scanner.Scan() {
		cards = append(cards, ParseCard(scanner.Text()))
	}
	return cards
}

// Score is calculated as: "The first match makes the card worth one point and
// each match after the first doubles the point value of that card."
// 0 winning number = score 0, 1 winning number = score 1, 2 numbers = 2, 3 numbers = 4, etc.
func (c Card) Score() int {
	winningNumbers := 0
	for _, ownedNumber := range c.OwnedNumbers {
		if utils.Contains(c.WinningNumbers, ownedNumber) {
			winningNumbers++
		}
	}

	if winningNumbers == 0 {
		return 0
	}

	// 2 pow 0 is 1, 2 pow 1 is 2, 2 pow 2 is 4, etc
	return utils.PowInt(2, winningNumbers-1)
}

func TotalCardScore(r io.Reader) int {
	cards := ParseCards(r)
	totalScore := 0
	for _, card := range cards {
		totalScore += card.Score()
	}
	return totalScore
}
