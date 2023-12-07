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
	// TODO: extract io.Reader to string to utility.
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		panic(err)
	}
	partNumbers := FindPartNumbers(buf.String())
	return partNumbers.Sum()
}

type Gear struct {
	ratios  []int
	gearPos utils.Coordinate
}

func FindGearRatios(schematic string) []int {
	grid := utils.NewGrid(schematic)
	currentNumber := ""

	// like above, find the numbers with only a * as neighbour
	// remember x/y of *
	// if two numbers have the same *, they are connected by a gear and should be added to the result.
	// TODO: Stuck, how do we determine when we have moved to a new number?
	//  Use * position as key in a map; just append whole numbers to the array of gear ratios

	gearRatios := make(map[utils.Coordinate][]int)

	currentGearPosition := utils.Coordinate{}

	grid.Iterate(func(pos utils.Coordinate, value rune) bool {
		if unicode.IsDigit(value) {
			currentNumber += string(value)

			grid.IterateNeighbours(pos, func(neighbourPos utils.Coordinate, neighbour rune) bool {
				// we found a gear; remember its position for this number
				if neighbour == '*' {
					currentGearPosition = neighbourPos
					return true
				}
				return false
			})
		} else {
			// we found the end of the current number, append it to the gear ratio map if it has a gear as neighbour
			if len(currentNumber) > 0 && currentGearPosition != (utils.Coordinate{}) {
				// do we have an entry for this gear yet?
				if _, ok := gearRatios[currentGearPosition]; !ok {
					gearRatios[currentGearPosition] = make([]int, 0)
				}
				gearRatios[currentGearPosition] = append(gearRatios[currentGearPosition], utils.ParseInt(currentNumber))
				// reset the current gear position
				currentGearPosition = utils.Coordinate{}
			}

			// reset, we're not on a number anymore.
			currentNumber = ""
		}
		return false
	})

	// get the gears with two ratios and multiply them
	result := make([]int, 0)
	for _, ratios := range gearRatios {
		if len(ratios) != 2 {
			continue
		}
		result = append(result, ratios[0]*ratios[1])
	}

	return result
}

func GetGearRatioSum(r io.Reader) int {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		panic(err)
	}
	ratios := FindGearRatios(buf.String())

	sum := 0
	for _, ratio := range ratios {
		sum += ratio
	}

	return sum
}
