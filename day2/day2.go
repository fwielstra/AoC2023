package day2

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/fwielstra/AoC2023/utils"
)

// TODO: look at https://medium.com/@ozoniuss/optimizing-go-string-operations-with-practical-examples-83df39b776fb for optimizations;
// examples:
//  - use strings.Index and split the string using subslices, saving allocations because it doesn't create new arrays (which strings.Split does)
//  - same with getting indivdual games, cutting bits off the remaining stringNa
//  - use strings.Cut to split a string in two parts, basically a utility for the above

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

// Power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
func (c CubeSet) Power() int {
	return c.Red * c.Green * c.Blue
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
	parts := utils.TrimmedFields(input, ',')

	for _, p := range parts {
		components := strings.Fields(p)
		number := utils.ParseInt(components[0])

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
	fragments := utils.TrimmedFields(input, ':')
	id := utils.ParseInt(strings.Fields(fragments[0])[1])
	roundStrings := utils.TrimmedFields(fragments[1], ';')
	rounds := make([]CubeSet, len(roundStrings))
	for i, r := range roundStrings {
		rounds[i] = ParseCubeSet(r)
	}

	return Game{
		id,
		rounds,
	}
}

// IsPossible returns true if the number of cubes in any round does not exceed the values in the given cube set
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

func (g Game) FewestPossible() CubeSet {
	result := CubeSet{}
	for _, round := range g.rounds {
		// new golang max builtin, woo
		result.Red = max(result.Red, round.Red)
		result.Green = max(result.Green, round.Green)
		result.Blue = max(result.Blue, round.Blue)
	}
	return result
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

func (g Games) SumOfFewestCubesPowered() int {
	result := 0
	for _, game := range g.games {
		fewestCubes := game.FewestPossible()
		result += fewestCubes.Power()
	}
	return result
}
