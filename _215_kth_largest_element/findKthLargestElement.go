package _215_kth_largest_element

func findKthLargest(nums []int, k int) int {
	l, r, n := 0, len(nums)-1, len(nums)
	if n < k {
		return -1
	}

	pivotPos := 0
	for l < r {
		pivotPos = getPartition(nums, l, r)
		if pivotPos == k {
			return nums[pivotPos]
		} else if pivotPos < k {
			l = pivotPos + 1
		} else {
			r = pivotPos - 1
		}
	}

	return -1
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
