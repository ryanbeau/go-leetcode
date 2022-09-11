package solution

type trie struct {
	children []*trie
	word     string
}

func (t *trie) Insert(word string) {
	end := t
	for i := 0; i < len(word); i++ {
		if end.children == nil {
			end.children = make([]*trie, 26)
		}
		c := word[i] - 'a'
		if end.children[c] == nil {
			end.children[c] = &trie{}
		}
		end = end.children[c]
	}
	end.word = word
}

func findNext(found *[]string, board [][]byte, t *trie, r int, c int) {
	const placeholder = '!'
	if board[r][c] == placeholder || t.children == nil {
		return
	}

	next := t.children[board[r][c]-'a']
	if next == nil {
		return
	}
	if next.word != "" {
		*found = append(*found, next.word)
		next.word = "" //prevent duplicate results - the optimal solution would be to remove from the Trie
	}

	temp := board[r][c]
	board[r][c] = placeholder

	if r < len(board)-1 {
		findNext(found, board, next, r+1, c)
	}
	if r > 0 {
		findNext(found, board, next, r-1, c)
	}
	if c < len(board[0])-1 {
		findNext(found, board, next, r, c+1)
	}
	if c > 0 {
		findNext(found, board, next, r, c-1)
	}

	board[r][c] = temp
}

func findWords(board [][]byte, words []string) (results []string) {
	root := &trie{}
	for _, w := range words {
		root.Insert(w)
	}

	w, h := len(board[0]), len(board)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if root.children != nil && root.children[board[r][c]-'a'] != nil {
				findNext(&results, board, root, r, c)
			}
		}
	}
	return
}
