package _000_common_sort_algorithm

func SelectionSort(nums []int) {
	n, mid := len(nums), 0
	if n < 1 {
		return
	}

	for i := 0; i < n; i++ {
		mid = i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[mid] {
				mid = j
			}
		}
		nums[i], nums[mid] = nums[mid], nums[i]
	}
}
