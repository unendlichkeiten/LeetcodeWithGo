package _000_common_sort_algorithm

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	quickSort(nums, 0, len(nums)-1)
	return
}

func quickSort(nums []int, l, r int) {
	pivotPos := 0
	if l < r {
		pivotPos = getPartition(nums, l, r)
		quickSort(nums, l, pivotPos-1)
		quickSort(nums, pivotPos+1, r)
	}
}

func getPartition(nums []int, l, r int) int {
	pivot := nums[l]
	for l < r {
		for l < r && pivot <= nums[r] {
			r--
		}
		nums[l] = nums[r]

		for l < r && nums[l] <= pivot {
			l++
		}
		nums[r] = nums[l]
	}

	nums[l] = pivot

	return l
}
