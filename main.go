package main

import (
	"fmt"
	"github.com/fwielstra/AoC2023/day3"
	"github.com/fwielstra/AoC2023/day4"
	"io"
	"os"
	"time"

	"github.com/fwielstra/AoC2023/day1"
	"github.com/fwielstra/AoC2023/day2"
)

func main() {
	day := os.Args[1]
	filename := os.Args[2]
	args := os.Args[3:]

	switch day {
	case "day1":
		runDay1(filename, args)
	case "day2":
		runDay2(filename, args)
	case "day3":
		runDay3(filename, args)
	case "day4":
		runDay4(filename, args)
	default:
		fmt.Printf("unrecognized day %s\n", day)
	}
}

func withFile(filename string, callback func(r io.Reader)) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s: %s\n", filename, err)
		os.Exit(1)
	}
	start := time.Now()
	callback(file)
	elapsed := time.Since(start)
	fmt.Printf("processing file %s took %s\n", filename, elapsed)

	err = file.Close()
	if err != nil {
		fmt.Printf("error closing file %s: %s\n", filename, err)
		os.Exit(1)
	}
}

func runDay1(filename string, args []string) {
	withFile(filename, func(r io.Reader) {
		result := day1.GetCalibrationValue(r)
		fmt.Printf("Result day 1 part 1 with input file %s: %d\n", filename, result)
	})

	withFile(filename, func(r io.Reader) {
		result := day1.GetCalibrationValueWords(r)
		fmt.Printf("Result day 1 part 2 with input file %s: %d\n", filename, result)
	})
}

func runDay2(filename string, args []string) {
	withFile(filename, func(r io.Reader) {
		games := day2.ParseGames(r)
		result := games.SumOfPossibleIds(day2.CubeSet{Red: 12, Green: 13, Blue: 14})
		fmt.Printf("Result day 2 part 1 with input file %s: %d\n", filename, result)
	})

	withFile(filename, func(r io.Reader) {
		games := day2.ParseGames(r)
		result := games.SumOfFewestCubesPowered()
		fmt.Printf("Result day 2 part 2 with input file %s: %d\n", filename, result)
	})
}

func runDay3(filename string, args []string) {
	withFile(filename, func(r io.Reader) {
		sum := day3.GetPartNumberSum(r)
		fmt.Printf("Result day 3 part 1 with input file %s: %d\n", filename, sum)
	})

	withFile(filename, func(r io.Reader) {
		sum := day3.GetGearRatioSum(r)
		fmt.Printf("Result day 3 part 2 with input file %s: %d\n", filename, sum)
	})
}

func runDay4(filename string, args []string) {
	withFile(filename, func(r io.Reader) {
		score := day4.TotalCardScore(r)
		fmt.Printf("Result day 4 part 1 with input file %s: %d\n", filename, score)
	})

	withFile(filename, func(r io.Reader) {
		result := day4.CountWonScratchCards(r)
		fmt.Printf("Result day 4 part 1 with input file %s: %d\n", filename, result)
	})
}
