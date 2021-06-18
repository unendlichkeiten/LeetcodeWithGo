package _075_sort_colors

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	pivotPos := 0
	if left < right {
		pivotPos = partition(nums, left, right)
		quickSort(nums, left, pivotPos-1)
		quickSort(nums, pivotPos+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && pivot <= nums[right] {
			right--
		}
		nums[left] = nums[right]

		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot
	return left
}
