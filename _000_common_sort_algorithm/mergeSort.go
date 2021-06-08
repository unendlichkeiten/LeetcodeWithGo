package _000_common_sort_algorithm

func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) {
	if l < r {
		mid := l + (r-l)/2
		mergeSort(nums, l, mid)
		mergeSort(nums, mid+1, r)
		merge(nums, l, mid, r)
	}
}

func merge(nums []int, l, mid, r int) {
	tmp := make([]int, 0)
	for _, v := range nums {
		tmp = append(tmp, v)
	}

	i, j, k := 0, 0, 0
	for i, j, k = l, mid+1, l; i <= mid && j <= r; k++ {
		if tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}

	for i <= mid {
		nums[k] = tmp[i]
		i++
		k++
	}

	for j <= r {
		nums[k] = tmp[j]
		j++
		k++
	}
}
