package _004_median_of_two_sorted_array

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenNums1 := len(nums1)
	lenNums2 := len(nums2)

	i, j := 0, 0
	nums := make([]int, 0)
	for i < lenNums1 && j < lenNums2 {
		if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			nums = append(nums, nums2[j])
			j++
		} else {
			nums = append(nums, nums1[i])
			nums = append(nums, nums2[j])
			i++
			j++
		}
	}

	for i < lenNums1 {
		nums = append(nums, nums1[i])
		i++
	}

	for j < lenNums2 {
		nums = append(nums, nums2[j])
		j++
	}

	if len(nums)%2 == 0 {
		return (float64)(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2
	}

	return (float64)(nums[len(nums)/2])
}
