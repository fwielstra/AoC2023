package day1

import (
	"strings"
	"testing"
)

func TestGetCalibrationValue(t *testing.T) {

	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
`

	r := strings.NewReader(input)
	result := GetCalibrationValue(r)

	if result != 142 {
		t.Errorf("Expected 142, got %d", result)
	}
}

func TestGetCalibrationValueWords(t *testing.T) {
	input := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen
`

	r := strings.NewReader(input)
	result := GetCalibrationValueWords(r)

	if result != 281 {
		t.Errorf("Expected 281, got %d", result)
	}
}
