package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var dotsRe = regexp.MustCompile(`[-,.!?'"]`)

func Top10(inputString string) []string {
	inputString = dotsRe.ReplaceAllString(inputString, "")
	words := strings.Fields(strings.ToLower(inputString))
	counts := map[string]int{}

	if len(words) == 0 {
		return words
	}

	for _, value := range words {
		counts[value]++
	}

	keys := make([]string, 0, len(counts))

	for key := range counts {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return (counts[keys[i]] > counts[keys[j]]) ||
			(counts[keys[i]] == counts[keys[j]] && keys[i] < keys[j])
	})

	return keys[0:10]
}
