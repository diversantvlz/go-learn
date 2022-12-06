package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var dotsRe = regexp.MustCompile(`[,.!?'"]`)

func Top10(inputString string) []string {
	inputString = dotsRe.ReplaceAllString(inputString, "")
	words := strings.Fields(strings.ToLower(inputString))
	counts := map[string]int{}

	for _, value := range words {
		counts[value]++
	}
	sort.SliceStable(counts, func(i, j int) bool {
		return i < j
	})
	return nil
}
