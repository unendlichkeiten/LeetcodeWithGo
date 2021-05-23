package _455_assign_cookie

func findContentChildren(g []int, s []int) int {
	quickSort(g, 0, len(g)-1)
	quickSort(s, 0, len(s)-1)

	i, j, n := 0, 0, 0
	for {
		if i >= len(g) || j >= len(s) {
			break
		} else if i < len(g) && j < len(s) && g[i] <= s[j] {
			n++
			i++
			j++
			continue
		} else {
			j++
		}
	}

	return n
}

func quickSort(A []int, left, right int) {
	if left < right {
		pivotPos := partition(A, left, right)
		quickSort(A, left, pivotPos-1)
		quickSort(A, pivotPos+1, right)
	}
}

// find pivot pos
func partition(A []int, left, right int) int {
	pivot := A[left]

	for left < right {
		for left < right && A[right] >= pivot {
			right--
		}
		A[left] = A[right]
		for left < right && A[left] <= pivot {
			left++
		}
		A[right] = A[left]
	}

	A[left] = pivot
	return left
}
