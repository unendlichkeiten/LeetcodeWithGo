package _000_common_sort_algorithm

func InsertSort(nums []int) {
	// empty slice
	if len(nums) < 1 {
		return
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
