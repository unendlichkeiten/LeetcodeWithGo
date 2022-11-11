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
		if j == -1 /*回退到第一位*/ || haystack[i] == src[j] {
			i, j = i+1, j+1
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
	// i 表示子串下标，j 表示当前的 next[i] 所对应回退的步数
	next, i, j := make([]int, len(src)), 0, -1

	for i < len(src)-1 { // i 取值范围为 0 ~ len(src)-2 ,因为进入 if 判断后 i 会自增
		// next[0] 是固定值
		if i == 0 {
			next[0] = -1
		}

		// 这里 j 表示匹配位置的前一位，匹配成功指针后移
		if j == next[0] /*回退到第一位*/ || src[j] == src[i] {
			i, j = i+1, j+1
			next[i] = j // j 表示当前的 next[i] 所对应的索引位置
		} else {
			j = next[j] // 指针回退
		}
	}

	return next
}
