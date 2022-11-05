package _1668_maximum_repeating_substring

// error resolution
// 这个只能统计重复子串的个数
func maxRepeatingError(sequence string, word string) int {
	i, j, n := 0, 0, 0
	for i < len(sequence) {
		for i < len(sequence) && j < len(word) {
			if sequence[i] == word[j] {
				i++
				j++
			} else {
				i = i - j + 1
				j = 0
			}
		}

		if j == len(word) {
			n++
		}
		j = 0
	}
	return n
}

func maxRepeating(sequence string, word string) int {
	repeatNums := make([]int, 0)
	for p := 0; p < len(sequence); p++ {
		i, j, n := p, 0, 0
		for i < len(sequence) && j < len(word) {
			if sequence[i] == word[j] {
				i++
				j++
			} else {
				break
			}

			if j == len(word) {
				n++
				j = 0
			}
		}

		repeatNums = append(repeatNums, n)
	}

	maxRepeat := 0
	for _, num := range repeatNums {
		if maxRepeat < num {
			maxRepeat = num
		}
	}

	return maxRepeat
}
