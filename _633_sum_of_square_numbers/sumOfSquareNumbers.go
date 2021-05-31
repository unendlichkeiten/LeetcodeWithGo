package _633_sum_of_square_numbers

import "math"

func judgeSquareSum(c int) bool {
	cslice := make([]int, int(math.Ceil(math.Sqrt(float64(c)))))

	for i := 0; i < len(cslice); i++ {
		cslice[i] = i
	}

	left, right := 0, len(cslice)
	for left <= right {
		if left*left+right*right > c {
			right--
		} else if left*left+right*right < c {
			left++
		} else {
			return true
		}
	}

	return false
}
