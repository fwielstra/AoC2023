package day3

import (
	"reflect"
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
