package _028_sub_string_index

// 模式匹配
func strStr(haystack string, src string) int {
	i, j := 0, 0
	for i < len(haystack) && j < len(src) {
		if haystack[i] == src[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == len(src) {
		return i - j
	}

	return -1
}

// KMP 算法
func strStrV2(haystack string, src string) int {
	i, j, next := 0, 0, getNext(src)
	for i < len(haystack) && j < len(src) {
		if j == -1 || haystack[i] == src[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(src) {
		return i - j
	}

	return -1
}

func getNext(src string) []int {
	next := make([]int, len(src))
	next[0] = -1         // 如果第一位失配，直接匹配主串的下一位
	i, j := 0, -1        // i 表示子串下表，j 表示当前的 next[i] 所对应回退的步数
	for i < len(src)-1 { // i 取值范围为 0~len(needle)-2 ,因为进入 if 判断后 i 会自增
		if j == -1 /*第一位匹配失败*/ || src[j] == src[i] {
			// 上面的比较仍然是假定失配位的前一位，因为当条件满足时指针会向后移动
			i++
			j++ // j 表示当前的 next[i] 所对应的索引
			next[i] = j
		} else {
			j = next[j] // 指针回退
		}
	}

	return next
}
