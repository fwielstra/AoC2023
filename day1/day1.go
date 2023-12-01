package day1

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

const digits = "0123456789"

func GetCalibrationValue(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	result := 0
	for scanner.Scan() {
		next := strings.TrimSpace(scanner.Text())
		// find first and last number
		first := strings.IndexAny(next, digits)
		last := strings.LastIndexAny(next, digits)

		// concat result
		number := string(next[first]) + string(next[last])
		nr, _ := strconv.Atoi(number)

		result += nr
	}
	return result
}

func GetCalibrationValueWords(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	// TODO: Refactor to streaming to avoid iterating over the lines twice

	// note: test input has "sixteen" but that's a deception, input only has 1-9
	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	result := ""
	for scanner.Scan() {
		next := strings.TrimSpace(scanner.Text())
		// luck: this replaces an input like "oneight" correctly into "1ight"
		result += replacer.Replace(next) + "\n"
	}

	return GetCalibrationValue(strings.NewReader(result))
}
