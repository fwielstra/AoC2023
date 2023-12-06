package day3

import (
	"github.com/fwielstra/AoC2023/utils"
	"io"
	"strings"
	"unicode"
)

// TODO: I presume input will fit in memory still, but if, hypothetically, it didn't,
//  how would we be able to read it? Read per 3 lines I suppose,
//  or read a line, remember the indices of the symbols and use them for the next line.
// TODO: Currently going row by row, col by col; consider if this is the best iteration order
//  or if going columns first is closer to memory layout / faster

// Alternative: parse input, log x / ys of all numbers and symbols, then check that result? idk, seems duplicate.
// Alternative: read input as one long string; assume it's rectangular; remove newlines, then do something clever with mod. idk.

// nasty

type PartNumbers []int

func (p PartNumbers) Sum() int {
	result := 0
	for _, num := range p {
		result += num
	}
	return result
}

func FindPartNumbers(schematic string) PartNumbers {
	currentNumber := ""
	isPart := false
	result := make(PartNumbers, 0)
	grid := utils.NewGrid(schematic)

	grid.Iterate(func(pos utils.Coordinate, value rune) bool {
		if unicode.IsDigit(value) {
			currentNumber += string(value)

			grid.IterateNeighbours(pos, func(_ utils.Coordinate, neighbour rune) bool {
				if !unicode.IsDigit(neighbour) && neighbour != '.' {
					isPart = true
					return true
				}
				return false
			})
		} else {
			if len(currentNumber) > 0 && isPart {
				result = append(result, utils.ParseInt(currentNumber))
			}

			// reset, we're not on a number anymore.
			currentNumber = ""
			isPart = false
		}
		return false
	})

	return result
}

func GetPartNumberSum(r io.Reader) int {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		panic(err)
	}
	partNumbers := FindPartNumbers(buf.String())
	return partNumbers.Sum()
}
