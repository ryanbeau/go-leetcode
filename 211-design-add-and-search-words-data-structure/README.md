# Problem
[Leetcode - 211. Design Add and Search Words Data Structure (Medium)](https://leetcode.com/problems/design-add-and-search-words-data-structure/)

## Takeaway
- With the constraints of lowercase english letters we can safely assume only `'a'` to `'z'`.
  > `word` in `addWord` consists of lowercase English letters.
- Recursive `Search` method would be easiest but we really only need recursion with the `'.'` wildcard characters.

## Solution
- Going with a hybrid `Search` method with recursion when the `'.'` wildcard is found and iterative for remaining letters.
- Within the Trie we can store a 26 length array of the children ( `'a'` to `'z'`).
- The `AddWord` method is near identical to the [208. Implement Trie - Solution](208-implement-trie-prefix-tree).
- Wildcard `'.'` letters will require iteration through the 26 letter array of children for any non `nil` WordDictionary.

An example of an entirely recursive solution to `Search`:
```go
func (t *WordDictionary) Search(word string) bool {
	if word == "" {
		return t.word
	}
	if t.children == nil {
		return false
	}

	//wildcard Search
	if word[0] == '.' {
		next := word[1:]
		for i := 0; i < len(t.children); i++ {
			if t.children[i] != nil && t.children[i].Search(next) {
				return true
			}
		}
		return false
	}

	//letter Search
	c := word[0] - 'a'
	return t.children[c] != nil && t.children[c].Search(word[1:])
}
```
