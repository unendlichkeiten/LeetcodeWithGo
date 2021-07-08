package _079_word_search

func exist(board [][]byte, word string) bool {
	i, j, k := 0, 0, 0

	visited := make([][]bool, len(board))
	for i, _ := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i < len(board) {
		for j < len(board[0]) {
			if check(board, visited, word, i, j, k) {
				return true
			}
		}
	}

	return false
}

func check(board [][]byte, visited [][]bool, word string, i, j, k int) bool {
	// 越界
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || k < 0 {
		return false
	}

	if board[i][j] == word[k] && k == len(word)-1 {
		return true
	}

	if visited[i][j] {
		return false
	} else if board[i][j] != word[k] {
		visited[i][j] = false
		return false
	} else {
		visited[i][j] = true
		return check(board, visited, word, i+1, j, k+1) ||
			check(board, visited, word, i-1, j, k+1) ||
			check(board, visited, word, i, j+1, k+1) ||
			check(board, visited, word, i, j-1, k+1)
	}
}
