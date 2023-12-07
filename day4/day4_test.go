package day4

import (
	"reflect"
	"strings"
	"testing"
)

const input = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func Test_ParseCard(t *testing.T) {
	result := ParseCard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	expectedWinningNumbers := []int{41, 48, 83, 86, 17}
	expectedOwnedNumbers := []int{83, 86, 6, 31, 17, 9, 48, 53}

	if !reflect.DeepEqual(result.WinningNumbers, expectedWinningNumbers) {
		t.Errorf("ParseCard() winning numbrs does not match, expected %+v, got %+v", expectedWinningNumbers, result.WinningNumbers)
	}

	if !reflect.DeepEqual(result.OwnedNumbers, expectedOwnedNumbers) {
		t.Errorf("ParseCard() owned numbers does not match, expected %+v, got %+v", expectedOwnedNumbers, result.OwnedNumbers)
	}
}

func Test_Score(t *testing.T) {
	cards := ParseCards(strings.NewReader(input))
	expectedScores := []int{8, 2, 2, 1, 0, 0}

	for i, card := range cards {
		if card.Score() != expectedScores[i] {
			t.Errorf("Expected score %d for card %d, got %d", expectedScores[i], i, card.Score())
		}
	}
}

func Test_TotalScore(t *testing.T) {
	result := TotalCardScore(strings.NewReader(input))
	if result != 13 {
		t.Errorf("TotalCardScore(): Expected total score of 13, got %d", result)
	}
}

func Test_CountWonScratchCards(t *testing.T) {
	result := CountWonScratchCards(strings.NewReader(input))
	if result != 30 {
		t.Errorf("CountWonScratchCards(): Expected total score of 13, got %d", result)
	}
}
