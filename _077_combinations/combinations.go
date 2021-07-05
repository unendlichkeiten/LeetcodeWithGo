package _077_combinations

func combine(n int, k int) [][]int {
	res, track := make([][]int, 0), make([]int, 0)

	backtrace(1, n, k, track, &res)
	return res
}

func backtrace(cur, n, k int, track []int, res *[][]int) {
	if len(track)+(n-cur+1) < k {
		return
	}

	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	track = append(track, cur)
	backtrace(cur+1, n, k, track, res)
	track = track[:len(track)-1]
	backtrace(cur+1, n, k, track, res)
}
