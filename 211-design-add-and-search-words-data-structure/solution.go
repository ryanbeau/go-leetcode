package solution

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

type WordDictionary struct {
	children []*WordDictionary
	word     bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (t *WordDictionary) AddWord(word string) {
	end := t
	for i := 0; i < len(word); i++ {
		if end.children == nil {
			end.children = make([]*WordDictionary, 26)
		}
		c := word[i] - 'a'
		if end.children[c] == nil {
			end.children[c] = &WordDictionary{}
		}
		end = end.children[c]
	}
	end.word = true
}

func (t *WordDictionary) Search(word string) bool {
	curr := t
	for i := 0; i < len(word); i++ {
		if curr.children == nil {
			return false
		}

		//wildcard Search - recursive
		if word[i] == '.' {
			for k := 0; k < len(curr.children); k++ {
				if curr.children[k] != nil && curr.children[k].Search(word[i+1:]) {
					return true
				}
			}
			return false
		}

		//letter Search - iterative
		curr = curr.children[word[i]-'a']
		if curr == nil {
			return false
		}
	}
	return curr.word
}
