package solution

func existNext(board [][]byte, word string, r int, c int, i int) bool {
	if board[r][c] != word[i] {
		return false
	}
	if i+1 == len(word) {
		return true
	}

	temp := board[r][c]
	board[r][c] = '!'

	result := r < len(board)-1 && existNext(board, word, r+1, c, i+1) ||
		r > 0 && existNext(board, word, r-1, c, i+1) ||
		c < len(board[0])-1 && existNext(board, word, r, c+1, i+1) ||
		c > 0 && existNext(board, word, r, c-1, i+1)

	board[r][c] = temp

	return result
}

func exist(board [][]byte, word string) bool {
	w, h := len(board[0]), len(board)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if existNext(board, word, r, c, 0) {
				return true
			}
		}
	}
	return false
}
