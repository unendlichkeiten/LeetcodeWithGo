package _036_sudoku

func IsValidSudoku(board [][]byte) bool {
	rows, columns := make([]map[int]int, len(board)), make([]map[int]int, len(board))
	subBoard := [3][3]map[int]int{}
	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] == '.' {
				continue
			}

			index := board[i][j] - '1'

			if len(rows[i]) == 0 {
				rows[i] = make(map[int]int)
			}
			if len(columns[j]) == 0 {
				columns[j] = make(map[int]int)
			}
			if len(subBoard[i/3][j/3]) == 0 {
				subBoard[i/3][j/3] = make(map[int]int)
			}

			rows[i][int(index)]++
			columns[j][int(index)]++
			subBoard[i/3][j/3][int(index)]++

			if rows[i][int(index)] > 1 || columns[j][int(index)] > 1 ||
				subBoard[i/3][j/3][int(index)] > 1 {
				return false
			}
		}
	}

	return true
}
