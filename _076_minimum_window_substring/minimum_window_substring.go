package _076_minimum_window_substring

import "math"

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	lenWin := math.MaxInt32
	ansL, ansR := -1, -1

	// 检查窗口是否包含所有元素
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		// 计算子串元素出现的频次
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		// 串口包含子串所有元素且 l <= r (移动左指针)
		for check() && l <= r {
			if r-l+1 < lenWin {
				// 记录窗口大小和位置
				lenWin = r - l + 1
				ansL, ansR = l, l+lenWin
			}
			// 减去子串元素再窗口出现的频次
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
