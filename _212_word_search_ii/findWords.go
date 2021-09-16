package _212_word_search_ii

import "LeetcodeWithGo/_079_word_search"

func FindWords(board [][]byte, words []string) []string {

	ans := make([]string, 0)
	for _, word := range words {
		if _079_word_search.Exist(board, word) {
			ans = append(ans, word)
		}
	}

	return ans

}
