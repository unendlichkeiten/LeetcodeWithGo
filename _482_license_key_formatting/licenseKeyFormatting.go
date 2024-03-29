package _482_license_key_formatting

/*
** 解题思路：
** 1. 计算字符串中字符数
** 2. 计算第一个分组字符数
** 3. 遍历原字符串，若为小写字符，转为大写，若为分隔符跳过
** 4. 将字符加入目标字符串中，每个分组结束追加字符‘-’
 */

func licenseKeyFormatting(s string, k int) string {
	// 计算字符串中字符数
	var cnt int
	for _, ch := range s {
		if ch != '-' {
			cnt++
		}
	}

	// 分组
	ans, n := make([]byte, 0), initN(cnt, k)
	for i, ch := range s {
		if ch == '-' {
			continue
		}

		if ch >= 'a' && ch <= 'z' {
			ch = ch - 32
		}

		if i > 0 && n == 0 {
			ans, n = append(ans, '-'), k
		}

		ans = append(ans, byte(ch))
		n--
	}

	return string(ans)
}

func initN(cnt, k int) int {
	if cnt%k == 0 {
		return k
	}

	return cnt % k
}
