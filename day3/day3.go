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
	rows := strings.Split(schematic, "\n")
	result := make(PartNumbers, 0)
	for y := range rows {
		col := rows[y]
		if col == "" {
			continue
		}

		for x := range col {
			value := rune(col[x])
			if unicode.IsDigit(value) {
				currentNumber += string(value)

				miny := max(0, y-1)
				maxy := min(len(rows)-1, y+1)

				// check neighbours; ensure we don't go out of bounds
				for suby := miny; suby <= maxy; suby++ {
					subrow := rows[suby]

					// make sure we get the length of the current subrow in case it's empty or different length from the current main row.
					minx := max(0, x-1)
					maxx := min(len(subrow)-1, x+1)

					for subx := minx; subx <= maxx; subx++ {
						neighbour := rows[suby][subx]
						if !unicode.IsDigit(rune(neighbour)) && neighbour != '.' {
							isPart = true
						}

					}
				}
			} else {
				if len(currentNumber) > 0 && isPart {
					result = append(result, utils.ParseInt(currentNumber))
				}

				// reset, we're not on a number anymore.
				currentNumber = ""
				isPart = false
			}
		}
	}

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
