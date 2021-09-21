package _058_length_of_last_word

func lengthOfLastWord(s string) int {
	start, end := 0, len(s)-1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			end--
			continue
		}
		break
	}

	for i := end; i >= 0; i-- {
		if s[i] == ' ' {
			start = i + 1
			break
		}
	}

	return end - start + 1
}
