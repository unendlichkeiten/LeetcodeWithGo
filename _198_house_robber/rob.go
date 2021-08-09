package _198_house_robber

func rob(nums []int) int {
	// dp[i] = max(dp[n-1], nums[i-1]+dp[i-2])
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	pre1, pre2, cur := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		cur = max(pre2+nums[i], pre1)
		pre2 = pre1
		pre1 = cur
	}

	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
