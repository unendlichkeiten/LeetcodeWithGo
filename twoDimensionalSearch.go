package leetCode

func TwoDimensionalSearch(array [][]int, rows, columns int, num int) bool {
	row, column := 0, columns-1
	for row < rows && column >= 0 {
		if array[row][column] > num {
			column--
		} else if array[row][column] < num {
			row++
		} else {
			return true
		}
	}

	return false
}
