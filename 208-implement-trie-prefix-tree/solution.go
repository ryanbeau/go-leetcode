package solution

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
	children []*Trie
	word     bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	end := t
	for i := 0; i < len(word); i++ {
		if end.children == nil {
			end.children = make([]*Trie, 26)
		}
		c := word[i] - 'a'
		if end.children[c] == nil {
			end.children[c] = &Trie{}
		}
		end = end.children[c]
	}
	end.word = true
}

func (t *Trie) find(s string) *Trie {
	end := t
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if end.children == nil || end.children[c] == nil {
			return nil
		}
		end = end.children[c]
	}
	return end
}

func (t *Trie) Search(word string) bool {
	if result := t.find(word); result != nil {
		return result.word
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	if result := t.find(prefix); result != nil {
		return true
	}
	return false
}
