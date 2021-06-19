package _069_sqrtx

func mySqrt(x int) int {
	l, r := 0, (x+1)/2
	sqrt := -1
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			sqrt = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return sqrt
}
