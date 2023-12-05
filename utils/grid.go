package utils

import "strings"

// TODO: unit tests?

type Grid struct {
	grid string
	// remember the split rows
	// TODO: figure out iterating over a grid with lookback/lookahead without having to split, e.g. using mod or whatever, to save memory.
	rows []string
}

func NewGrid(rawInput string) *Grid {
	return &Grid{
		grid: rawInput,
		rows: strings.Split(rawInput, "\n"),
	}
}

// Iterate iterates over the grid, rows (y) by cols (x), calling the callback with x, y and the current character.
// Have the callback return "true" to stop iteration.
func (g *Grid) Iterate(callback func(int, int, rune) bool) {
	for y := range g.rows {
		col := g.rows[y]

		for x := range col {
			value := rune(col[x])
			shouldBreak := callback(y, x, value)
			if shouldBreak {
				return
			}
		}
	}
}

// if the callback returns true, iteration will stop.

// IterateNeighbours will call the given callback for every x/y coordinate
// around the given x/y coordinate, including diagonals, excluding out of bounds values.
// NOTE: untested with uneven rows
// TODO: Add variant that excludes diagonals.
func (g *Grid) IterateNeighbours(y, x int, callback func(rune) bool) {
	miny := max(0, y-1)
	maxy := min(len(g.rows)-1, y+1)

	for suby := miny; suby <= maxy; suby++ {
		subrow := g.rows[suby]

		// make sure we get the length of the current subrow in case it's empty or different length from the current main row.
		minx := max(0, x-1)
		maxx := min(len(subrow)-1, x+1)

		for subx := minx; subx <= maxx; subx++ {
			neighbour := g.rows[suby][subx]
			callback(rune(neighbour))
		}
	}
}
