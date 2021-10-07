package _434_number_of_segments_in_a_string

func countSegments(s string) int {
	// remove front blank spaces
	index, cnt := 0, 0
	for _, ch := range s {
		if ch != ' ' {
			break
		}
		index++
	}

	if index == len(s) {
		return 0
	}

	s = s[index:]

	pre := int32(s[0])
	for i, ch := range s {
		if (ch == ' ' && pre != ' ') || (i == len(s)-1 && s[i] != ' ') {
			cnt++
		}
		pre = ch
	}

	return cnt
}
