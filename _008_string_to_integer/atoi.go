package _008_string_to_integer

import "math"

func myAtoi(s string) int {
	var s1 string
	if s1 = trimSpace(s); len(s1) == 0 {
		return 0
	}

	sign := 0
	// ensure string sign
	if s1[0] == '-' {
		sign = -1
		s1 = s1[1:]
	} else if s1[0] == '+' {
		s1 = s1[1:]
		sign = 1
	} else if s1[0] >= '0' && s1[0] <= '9' {
		sign = 1
	} else {
		return 0
	}

	var num float64
	for i := 0; i < len(s1) && s1[i] >= '0' && s1[i] <= '9'; i++ {
		num = num*10 + (float64)(s1[i]-'0')
	}

	num = float64(sign) * num
	if num > math.MaxInt32 {
		return math.MaxInt32
	} else if num < math.MinInt32 {
		return math.MinInt32
	}

	return int(num)
}

func trimSpace(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			s = s[i:]
			break
		}
	}
	return s
}
