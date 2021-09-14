package _524_delete_words

func FindLongestWord(s string, dictionary []string) string {
	var maxWord string

	for _, word := range dictionary {
		i := 0
		for j := range s {
			if s[j] == word[i] {
				i++

				if i == len(word) {
					if len(word) > len(maxWord) || len(word) == len(maxWord) && word < maxWord {
						maxWord = word
					}
					break
				}
			}
		}
	}

	return maxWord
}
