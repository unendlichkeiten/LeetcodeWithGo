package _452_minimum_number_of_arrows

func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return 1
	}

	quickSortSlice(points, 0, len(points))

	i, j, n := 1, 0, 0
	for i < len(points) {
		if points[i][0] > points[j][len(points[j])-1] {
			n++
			j = i
			i++
			continue
		}
		i++
	}

	return n
}

func quickSortSlice(points [][]int, left, right int) {
	if left < right {
		pivotPos := partitionSlice(points, left, right)
		quickSortSlice(points, left, pivotPos - 1)
		quickSortSlice(points, pivotPos + 1, right)
	}
}

func partitionSlice(points [][]int, left, right int) int {
	pivot := points[left]

	for left < right {
		for left < right && points[right][len(points)-1] >= pivot[len(pivot)-1] {
			right--
		}
		points[left] = points[right]
		for left < right && points[left][len(points)-1] <= pivot[len(pivot)-1] {
			left++
		}
		points[right] = points[left]
	}

	points[left] = pivot
	return left
}
