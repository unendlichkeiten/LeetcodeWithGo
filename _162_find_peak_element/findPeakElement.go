package _162_find_peak_element

func FindPeakElement(nums []int) int {
	maxIndex := 0
	for i, _ := range nums[:len(nums)-1] {
		if nums[i] < nums[i+1] {
			maxIndex = i + 1
			continue
		}
		return maxIndex
	}
	return maxIndex
}
