package _166_fraction_to_recurring_decimal

import (
	"math"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	// 整除
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	// 小数
	ans := make([]byte, 0)
	// 判定符号
	if numerator < 0 != (denominator < 0) {
		ans = append(ans, byte('-'))
	}
	// 计算整数部分
	numeratorAbs := int(math.Abs(float64(numerator)))
	denominatorAbs := int(math.Abs(float64(denominator)))
	integer := numeratorAbs / denominatorAbs
	ans = append(ans, []byte(strconv.Itoa(integer))...)

	// 小数点
	ans = append(ans, '.')

	remainder, remainderMap := numeratorAbs%denominatorAbs, make(map[int]int)
	// 余数不等于 0，且余数第一次出现
	for remainder != 0 && remainderMap[remainder] == 0 {
		// 记录余数第一次出现的位置
		remainderMap[remainder] = len(ans)
		remainder = 10 * remainder
		ans = append(ans, []byte(strconv.Itoa(remainder/denominatorAbs))...)
		remainder = remainder % denominatorAbs
	}

	// 循环小数
	if remainder > 0 {
		ans = append(ans, ')')
		loopStr := append([]byte{'('}, ans[remainderMap[remainder]:]...)
		ans = append(ans[:remainderMap[remainder]], loopStr...)
	}

	return string(ans)
}
