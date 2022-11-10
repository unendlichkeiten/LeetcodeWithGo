package _028_sub_string_index

// 模式匹配
func strStr(haystack string, needle string) int {
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == len(needle) {
		return i - j
	}

	return -1
}

// KMP 算法
func strStrV2(haystack string, needle string) int {
	return 0
}
