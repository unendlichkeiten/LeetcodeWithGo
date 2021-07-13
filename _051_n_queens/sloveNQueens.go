package _051_n_queens

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	backTrace(0, board, &res, n)

	return res
}

func backTrace(r int, board [][]byte, res *[][]string, n int) {
	// 每一行都放置一个皇后之后返回
	if r == n {
		temp := make([]string, 0)
		for _, row := range board {
			temp = append(temp, string(row))
		}
		*res = append(*res, temp)
		return
	}

	for c := 0; c < n; c++ {
		if isValid(r, c, board) {
			board[r][c] = 'Q'
			backTrace(r+1, board, res, n)
		}
		board[r][c] = '.'
	}

}

func isValid(r, c int, board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'Q' && (i == r || j == c || i+j == r+c || i-j == r-c) {
				return false
			}
		}
	}
	return true
}
