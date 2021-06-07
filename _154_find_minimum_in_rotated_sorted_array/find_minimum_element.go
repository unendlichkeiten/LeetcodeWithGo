package _154_find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	l, r, mid := 0, len(nums)-1, 0

	for l <= r {
		mid = l + (r-l)/2
		if nums[l] < nums[r] {
			return nums[l]
		}

		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			r--
		}
	}

	return nums[l]
}
