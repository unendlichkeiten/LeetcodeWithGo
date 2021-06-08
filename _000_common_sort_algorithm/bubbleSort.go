package _000_common_sort_algorithm

func BubbleSort(nums []int) {
	n := len(nums)
	if n < 1 {
		return
	}

	flag := false
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}
