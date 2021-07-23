package _047_permutations_ii

import (
	"sort"
)

func PermuteUnique(nums []int) [][]int {

	// nums sorting
	sort.Ints(nums)

	res, track, used := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	backtrace(nums, track, &res, &used)

	return res
}

func backtrace(nums, track []int, res *[][]int, used *[]bool) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)

		return
	}

	for i, num := range nums {
		if trimBranch(nums, i, *used) {
			continue
		}

		if (*used)[i] == false {
			(*used)[i] = true
			track = append(track, num)
			backtrace(nums, track, res, used)
			track = track[:len(track)-1]
			(*used)[i] = false
		}
	}
}

func trimBranch(nums []int, index int, used []bool) bool {
	if index > 0 &&
		nums[index] == nums[index-1] &&
		used[index-1] == false {
		return true
	}
	return false
}
