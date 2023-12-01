package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fwielstra/AoC2023/day1"
)

func main() {
	day := os.Args[1]
	filename := os.Args[2]
	args := os.Args[3:]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s: %s\n", filename, err)
		os.Exit(1)
	}

	switch day {
	case "day1":
		runDay1(filename, args)
	default:
		fmt.Printf("unrecognized day %s\n", day)
	}

	file.Close()
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
