package _34_find_first_last_position

// 使用二分法查找有序数组中出现的 target 首尾位置
func searchRange(nums []int, target int) []int {
	i, j, mid := 0, len(nums)-1, 0
	firstPost, lastPos := -1, -1

	// first position
	for i < j {
		mid = (i + j) / 2
		if nums[mid] >= target {
			j = mid
		} else {
			i = mid + 1
		}
	}

	if len(nums) == 0 || nums[i] != target {
		return []int{firstPost, lastPos}
	}
	firstPost = i

	// last position
	i, j = 0, len(nums)-1
	for i < j {
		mid = (i + j) / 2
		if nums[mid] >= target+1 {
			j = mid
		} else {
			i = mid + 1
		}
	}

	if nums[j] > target {
		lastPos = j - 1
	} else {
		lastPos = j
	}

	return []int{firstPost, lastPos}
}
