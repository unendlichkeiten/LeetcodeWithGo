package _46_permutations

func permute(nums []int) [][]int {
	res, track := make([][]int, 0), make([]int, 0)

	backtrace(nums, track, &res)
	return res
}

func backtrace(nums, track []int, res *[][]int) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for _, num := range nums {
		if numInTrack(track, num) {
			continue
		}

		// dfs
		track = append(track, num)
		backtrace(nums, track, res)

		// backtracking
		track = track[:len(track)-1]
	}
}

func numInTrack(track []int, num int) bool {
	for _, v := range track {
		if v == num {
			return true
		}
	}
	return false
}
