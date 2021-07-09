package _079_word_search

func exist(board [][]byte, word string) bool {

	visited := make([][]bool, len(board))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			for i, _ := range visited {
				visited[i] = make([]bool, len(board[0]))
			}
			if check(board, &visited, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func check(board [][]byte, visited *[][]bool, word string, i, j, k int) bool {
	// 越界处理
	if i < 0 || i >= len(board) ||
		j < 0 || j >= len(board[0]) || (*visited)[i][j] {
		return false
	}

	if (board[i][j]) != word[k] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	// 回溯时还原已访问的单元格
	defer func() {
		(*visited)[i][j] = false
	}()

	(*visited)[i][j] = true

	return check(board, visited, word, i+1, j, k+1) ||
		check(board, visited, word, i-1, j, k+1) ||
		check(board, visited, word, i, j+1, k+1) ||
		check(board, visited, word, i, j-1, k+1)
}
