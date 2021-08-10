package _413_arithmetic_slices

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	counts, sumCount := make([]int, len(nums)), 0
	// 初始差值
	diff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == diff {
			counts[i] = counts[i-1] + 1
			sumCount += counts[i]
		} else {
			// 重新计算差值
			diff = nums[i] - nums[i-1]
		}
	}

	return sumCount
}
