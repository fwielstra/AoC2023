package day3

import (
	"reflect"
	"strings"
	"testing"
)

const schematic = `467..114.0
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_FindPartNumbers(t *testing.T) {
	result := FindPartNumbers(schematic)
	expected := PartNumbers{467, 35, 633, 617, 592, 755, 664, 598}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Result does not match, expected %+v, got %+v", expected, result)
	}

	if result.Sum() != 4361 {
		t.Errorf("Expected sum of 4361, got %d", result.Sum())
	}
}

func Test_FindGearRatios(t *testing.T) {
	result := FindGearRatios(schematic)
	expected := []int{16345, 451490}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Result does not match, expected %+v, got %+v", expected, result)
	}
}

func Test_GetGearRatioSum(t *testing.T) {
	result := GetGearRatioSum(strings.NewReader(schematic))

	if result != 467835 {
		t.Errorf("GetGearRatioSum() result does not match, expected 467835, got %d", result)
	}
}
