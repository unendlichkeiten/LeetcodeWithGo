package _524_longest_word_in_dictionary

import (
	"sort"
)

func findLongestWord(s string, d []string) string {
	sort.SliceStable(d, func(i, j int) bool {
		// sort by string length reverse, if string length identical sort by char sort
		return len(d[i]) == len(d[j]) && d[i] < d[j] || len(d[i]) > len(d[j])
	})

	max := ""
	for _, v := range d {
		i, j := 0, 0
		for i < len(s) && j < len(v) {
			if s[i] != v[j] {
				i++
				continue
			}
			i, j = i+1, j+1
		}

		if j == len(v) && (len(v) > len(max) || max == "") {
			max = v
		}
	}

	return max
}
