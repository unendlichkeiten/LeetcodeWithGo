package _070_climbing_stairs

func ClimbStairs(n int) int {
	// n > 2
	// dp[n] = dp[n-1] + dp[n-2]
	if n < 0 {
		return 0
	}

	if n <= 2 {
		return n
	}

	return ClimbStairs(n-1) + ClimbStairs(n-2)
}

// ClimbStairsv2
func ClimbStairsv2(n int) int {
	// n > 2
	// dp[n] = dp[n-1] + dp[n-2]
	if n < 0 {
		return 0
	}

	if n <= 2 {
		return n
	}

	pre1, pre2, cur := 1, 2, 0
	for i := 3; i <= n; i++ {
		cur = pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}

	return cur
}
