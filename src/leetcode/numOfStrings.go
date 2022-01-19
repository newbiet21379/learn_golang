package leetcode

import "strings"

func numOfStrings(patterns []string, word string) int {
	var (
		count int
	)
	for _, s := range patterns {
		if strings.Contains(word, s) {
			count++
		}
	}
	return count
}
