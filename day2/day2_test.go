package day2

import (
	"fmt"
	"strings"
	"testing"
)

const testInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestParseCubeSet(t *testing.T) {
	tests := []struct {
		input    string
		expected CubeSet
	}{
		{
			input:    "1 red, 2 green, 3 blue",
			expected: CubeSet{Red: 1, Green: 2, Blue: 3},
		},
		// multiple digit numbers
		{
			input:    "10 red, 1337 green, 99999999999 blue",
			expected: CubeSet{Red: 10, Green: 1337, Blue: 99999999999},
		},
		// partial
		{
			input:    "1 red",
			expected: CubeSet{Red: 1, Green: 0, Blue: 0},
		},
		// empty?
		{
			input:    "",
			expected: CubeSet{Red: 0, Green: 0, Blue: 0},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("should return %+v for input '%s'", tt.expected, tt.input), func(t *testing.T) {
			result := ParseCubeSet(tt.input)
			if tt.expected != result {
				t.Errorf("ParseCubeSet() expected %+v, got %+v", tt.expected, result)
			}
		})
	}
}

func TestParseGame(t *testing.T) {
	tests := []struct {
		input    string
		expected Game
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: Game{id: 1, rounds: []CubeSet{
				{Blue: 3, Red: 4},
				{Red: 1, Green: 2, Blue: 6},
				{Green: 2},
			}},
		},
		{
			// high game number
			input: "Game 9999: 3 blue, 4 red",
			expected: Game{id: 9999, rounds: []CubeSet{
				{Blue: 3, Red: 4},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("should return %+v for input '%s'", tt.expected, tt.input), func(t *testing.T) {
			result := ParseGame(tt.input)
			if tt.expected.id != result.id {
				t.Errorf("ParseGame() expected ID %d, got %d", tt.expected.id, result.id)
			}

			if len(result.rounds) != len(tt.expected.rounds) {
				t.Errorf("ParseGame() expected %d rounds, got %d", len(tt.expected.rounds), len(result.rounds))
			}

			for i, round := range result.rounds {
				if round != tt.expected.rounds[i] {
					t.Errorf("ParseGame() expected round %+v, got %+v", tt.expected.rounds[i], round)
				}
			}
		})
	}
}

func TestParseGames(t *testing.T) {
	games := ParseGames(strings.NewReader(testInput))

	if len(games.games) != 5 {
		t.Errorf("expected %d games, got %d", 5, len(games.games))
	}

	if games.games[4].id != 5 {
		t.Errorf("expected id 5, got %d", games.games[4].id)
	}

	if len(games.games[0].rounds) != 3 {
		t.Errorf("expected game 1 to have 3 rounds, got %d", len(games.games[0].rounds))
	}
}

func TestIsPossible(t *testing.T) {

	tests := []struct {
		rounds   []CubeSet
		maxes    CubeSet
		expected bool
	}{
		{
			rounds:   []CubeSet{{Red: 1, Green: 1, Blue: 1}},
			maxes:    CubeSet{Red: 1, Green: 1, Blue: 1},
			expected: true,
		},
		// these are from the example
		{
			rounds:   []CubeSet{{Red: 4, Blue: 4}, {Red: 1, Green: 6, Blue: 2}, {Green: 2}},
			maxes:    CubeSet{Red: 12, Green: 13, Blue: 14},
			expected: true,
		},
		{
			rounds:   []CubeSet{{Blue: 1, Green: 3}, {Red: 1, Green: 3, Blue: 4}, {Green: 1, Blue: 1}},
			maxes:    CubeSet{Red: 12, Green: 13, Blue: 14},
			expected: true,
		},
		{
			rounds:   []CubeSet{{Red: 20, Blue: 6, Green: 8}, {Red: 4, Green: 13, Blue: 5}, {Red: 1, Green: 5, Blue: 1}},
			maxes:    CubeSet{Red: 12, Green: 13, Blue: 14},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("IsPossible() should return %t for input '%+v'", tt.expected, tt.rounds), func(t *testing.T) {
			game := Game{id: 1, rounds: tt.rounds}
			result := game.IsPossible(tt.maxes)

			if tt.expected != result {
				t.Errorf("IsPossible() expected %t for game %+v and maxes %+v, got %+v", tt.expected, game, tt.maxes, result)
			}
		})
	}
}

func TestIdsOfPossible(t *testing.T) {
	games := ParseGames(strings.NewReader(testInput))

	ids := games.PossibleIds(CubeSet{Red: 12, Green: 13, Blue: 14})

	if len(ids) != 3 || ids[0] != 1 || ids[1] != 2 || ids[2] != 5 {
		t.Errorf("Expected ids 1, 2 and 5, got %+v", ids)
	}
}

func TestSumOfPossible(t *testing.T) {
	games := ParseGames(strings.NewReader(testInput))

	result := games.SumOfPossibleIds(CubeSet{Red: 12, Green: 13, Blue: 14})

	if result != 8 {
		t.Errorf("Expected sum 8, got %+d", result)
	}
}
