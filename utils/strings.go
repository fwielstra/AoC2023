package utils

import (
	"strconv"
	"strings"
)

func SplitAndTrim(s string, sep string) []string {
	result := strings.Split(s, sep)
	for i := range result {
		result[i] = strings.TrimSpace(result[i])
	}
	return result
}

func TrimmedFields(s string, sep rune) []string {
	result := strings.FieldsFunc(s, func(r rune) bool {
		return r == sep
	})
	for i := range result {
		result[i] = strings.TrimSpace(result[i])
	}
	return result
}

// TrimmedIntFields splits on space, trims, and parses result as number
func TrimmedIntFields(s string) []int {
	fields := strings.Fields(s)
	result := make([]int, len(fields))
	for i, value := range fields {
		result[i] = ParseInt(strings.TrimSpace(value))
	}
	return result
}

// ParseInt parses integers, panics if there's an error
func ParseInt(s string) int {
	result, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return result
}
