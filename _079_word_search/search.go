package _079_word_search

// 这个标准答案
type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func Exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}

	for i, row := range board {
		for j := range row {
			if Check(board, word, vis, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func Check(board [][]byte, word string, vis [][]bool, i, j, k int) bool {
	if board[i][j] != word[k] { // 剪枝：当前字符不匹配
		return false
	}
	if k == len(word)-1 { // 单词存在于网格中
		return true
	}
	vis[i][j] = true
	defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
	for _, dir := range directions {
		if newI, newJ := i+dir.x, j+dir.y; 0 <= newI &&
			newI < len(board) &&
			0 <= newJ &&
			newJ < len(board[0]) &&
			!vis[newI][newJ] {
			if Check(board, word, vis, newI, newJ, k+1) {
				return true
			}
		}
	}
	return false
}
