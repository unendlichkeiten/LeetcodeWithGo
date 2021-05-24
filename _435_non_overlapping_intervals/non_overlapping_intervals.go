package _435_non_overlapping_intervals

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	quickSortSlice(intervals, 0, len(intervals)-1)

	i, j, n := 1, 0, 0 // i, index of intervals
	for i < len(intervals) {
		if intervals[i][0] < intervals[j][len(intervals[j])-1] {
			n++
			continue
		}
		j = i
	}

	return n
}

func quickSortSlice(intervals [][]int, left, right int) {
	if left < right {
		pivotPos := partitionSlice(intervals, left, right)
		quickSortSlice(intervals, left, pivotPos-1)
		quickSortSlice(intervals, pivotPos+1, right)
	}
}

func partitionSlice(intervals [][]int, left, right int) int {
	pivot := intervals[left]

	for left < right {
		for left < right && intervals[right][len(intervals[right])-1] >= pivot[len(pivot)-1] {
			right--
		}
		intervals[left] = intervals[right]
		for left < right && intervals[left][len(intervals[left])-1] <= pivot[len(pivot)-1] {
			left++
		}
		intervals[right] = intervals[left]
	}

	intervals[left] = pivot
	return left
}
