package day4

import (
	"bufio"
	"github.com/fwielstra/AoC2023/utils"
	"io"
	"strings"
)

type Card struct {
	idx            int
	WinningNumbers []int
	OwnedNumbers   []int
	// remember the number of matches for multiple calls
	matches int
}

func ParseCard(input string) Card {
	prefix, values, _ := strings.Cut(input, ":")
	_, id, _ := strings.Cut(prefix, " ")
	winning, owned, _ := strings.Cut(values, "|")

	// idx is the card's position in the original array, so we can index on idx + score for day 2.
	result := Card{
		idx:            utils.ParseInt(id) - 1,
		WinningNumbers: utils.TrimmedIntFields(winning),
		OwnedNumbers:   utils.TrimmedIntFields(owned),
	}

	// Precalculate the number of matches
	result.matches = result.calculateMatches()

	return result
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

func (c Card) calculateMatches() int {
	matches := 0
	for _, ownedNumber := range c.OwnedNumbers {
		if utils.Contains(c.WinningNumbers, ownedNumber) {
			matches++
		}
	}
	return matches
}

// Matches returns the number of matches / winning numbers
func (c Card) Matches() int {
	return c.matches
}

// Score is calculated as: "The first match makes the card worth one point and
// each match after the first doubles the point value of that card."
// 0 winning number = score 0, 1 winning number = score 1, 2 numbers = 2, 3 numbers = 4, etc.
func (c Card) Score() int {
	winningNumbers := c.Matches()
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

type Cards []Card

func CountWonScratchCards(r io.Reader) int {
	cards := ParseCards(r)
	wonCards := len(cards)

	// Add all cards and all won cards to a queue of cards yet to be checked
	cardQueue := make([]Card, len(cards))
	copy(cardQueue, cards)

	for len(cardQueue) != 0 {
		// pop from queue; if cardQueue was in the same scope we could do it in one line.
		current := cardQueue[0]
		cardQueue = cardQueue[1:]
		score := current.Matches()
		if score > 0 {
			startIdx := current.idx + 1
			endIdx := startIdx + score
			toAppend := cards[startIdx:endIdx]
			wonCards += len(toAppend)
			cardQueue = append(cardQueue, toAppend...)
		}
	}

	return wonCards
}
