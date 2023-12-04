package day2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/fwielstra/AoC2023/utils"
)

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

// ParseCubeSet expects a CSV of format "3 Blue, 4 Red, 2 Green"
// input may not have all colors, may have multiple numbers
// TODO: We might be able to parse the whole thing in a single regex instead of split and multiple iterations
func ParseCubeSet(input string) CubeSet {
	result := CubeSet{}

	// check for empty lines to be sure
	if input == "" {
		fmt.Println("empty line")
		return result
	}

	// split and trim
	parts := utils.SplitAndTrim(input, ",")

	for _, p := range parts {
		components := utils.SplitAndTrim(p, " ")
		number, err := strconv.Atoi(components[0])

		if err != nil {
			panic(err)
		}

		// can remove the switch if we use a map instead but eh.
		switch components[1] {
		case "red":
			result.Red = number
		case "green":
			result.Green = number
		case "blue":
			result.Blue = number
		}
	}

	return result
}

type Game struct {
	id     int
	rounds []CubeSet
}

func ParseGame(input string) Game {
	// get game ID
	// Game 1: 3 Blue, 4 Red; 1 Red, 2 Green, 6 Blue; 2 Green
	fragments := utils.SplitAndTrim(input, ": ")
	id, _ := strconv.Atoi(utils.SplitAndTrim(fragments[0], " ")[1])
	roundStrings := utils.SplitAndTrim(fragments[1], ";")
	rounds := make([]CubeSet, len(roundStrings))
	for i, r := range roundStrings {
		rounds[i] = ParseCubeSet(r)
	}

	return Game{
		id,
		rounds,
	}
}

// A set of games is only possible if the number of cubes in any round
// exceeds the given max
func (g Game) IsPossible(maxes CubeSet) bool {
	for _, r := range g.rounds {
		isPossible := r.Red <= maxes.Red && r.Green <= maxes.Green && r.Blue <= maxes.Blue
		// early return if possible.
		if !isPossible {
			return false
		}
	}
	return true
}

type Games struct {
	games []Game
}

func ParseGames(input io.Reader) Games {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	games := make([]Game, 0)
	for scanner.Scan() {
		game := ParseGame(scanner.Text())
		games = append(games, game)
	}

	return Games{
		games,
	}
}

func (g Games) PossibleIds(maxes CubeSet) []int {
	result := make([]int, 0)
	for _, game := range g.games {
		if game.IsPossible(maxes) {
			result = append(result, game.id)
		}
	}
	return result
}

func (g Games) SumOfPossibleIds(maxes CubeSet) int {
	result := 0
	for _, possibleId := range g.PossibleIds(maxes) {
		result += possibleId
	}
	return result
}
