package _007_reverse_integer

import "math"

func reverse(x int) int {
	rev := 0
	for x != 0 {
		rev = x%10 + rev*10
		x /= 10
	}

	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}

	return rev
}
