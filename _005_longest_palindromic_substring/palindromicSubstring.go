package _005_longest_palindromic_substring

// 步骤：
// 1. 设置最长子串窗口大小为 len(s)
// 2. 判断该子串是否为回文子串
// 3. 若为 true 返回子串
// 4. 若为 false 向右移动子串窗口位置，判断子串是否为回文子串
// 5. 子串窗口大小 -1 ，重复执行 2 - 4 步骤。

func longestPalindrome(s string) string {
	max := len(s)

	for max > 0 {
		i, j := 0, max-1
		for i >= 0 && j < len(s) {
			if isPalindromic(s[i : j+1]) {
				return s[i : j+1]
			}
			i++
			j++
		}
		max--
	}

	return ""
}

func isPalindromic(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
