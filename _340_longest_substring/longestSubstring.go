package _340_longest_substring

type strMap struct {
	smap map[int32]int
	size int
}

func lengthOfLongestSubstringKDistinct(s string, k int) string {
	strStr := &strMap{
		smap: make(map[int32]int),
		size: 0,
	}

	// i 数组左指针
	// j 数组右指针
	i, j := 0, 0
	max_i, max_j := 0, 0
	for _, ch := range s {
		if _, ok := strStr.smap[ch]; !ok {
			strStr.size++
		}

		for strStr.size > k {
			if j-i > max_j-max_i {
				max_j = j
				max_i = i
			}
			strStr.smap[int32(s[i])]--
			if strStr.smap[int32(s[i])] == 0 {
				strStr.size--
			}
			i++
		}

		strStr.smap[ch]++
		j++
	}

	return s[max_i:max_j]
}
