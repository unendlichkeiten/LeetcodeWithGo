package _088_merge_sorted_array

// merge nums2 into nums1 and keep nums1 ordered
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < m+n {
		return
	}

	// process nums1 and nums2 sequence order
	nums1Right, nums2Right, mergeRight := m-1, n-1, m+n-1
	for nums1Right >= 0 && nums2Right >= 0 {
		if nums1[nums1Right] < nums2[nums2Right] {
			nums1[mergeRight] = nums2[nums2Right]
			nums2Right--
		} else {
			nums1[mergeRight] = nums1[nums1Right]
			nums1Right--
		}
		mergeRight--
	}

	for mergeRight >= 0 && nums2Right >= 0 {
		nums1[mergeRight] = nums2[nums2Right]
		nums2Right--
		mergeRight--
	}
}
