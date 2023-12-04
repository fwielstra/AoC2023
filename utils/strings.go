package utils

import "strings"

func SplitAndTrim(s string, sep string) []string {
	result := strings.Split(s, sep)
	for i := range result {
		result[i] = strings.TrimSpace(result[i])
	}
	return result
}
