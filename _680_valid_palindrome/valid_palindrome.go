package _680_valid_palindrome

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			subLeft, subRight := true, true
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					subLeft = false
					break
				}
			}

			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					subRight = false
					break
				}
			}

			return subLeft || subRight
		}
	}

	return true
}
