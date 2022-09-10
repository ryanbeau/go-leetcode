package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type trieFN func(t *testing.T, trie *Trie)

func insert(word string) trieFN {
	return func(t *testing.T, trie *Trie) {
		trie.Insert(word)
	}
}

func assertSearch(word string, expected bool) trieFN {
	return func(t *testing.T, trie *Trie) {
		actual := trie.Search(word)
		assert.Equal(t, expected, actual, "Search(\"%s\")", word)
	}
}

func assertStartsWith(prefix string, expected bool) trieFN {
	return func(t *testing.T, trie *Trie) {
		actual := trie.StartsWith(prefix)
		assert.Equal(t, expected, actual, "StartsWith(\"%s\")", prefix)
	}
}

func TestTrie(t *testing.T) {
	tests := [][]trieFN{
		{
			insert("apple"),
			assertSearch("apple", true),
			assertSearch("app", false),
			assertStartsWith("app", true),
			insert("app"),
			assertSearch("app", true),
		},
		{
			assertSearch("nothing", false),
			assertStartsWith("empty", false),
		},
	}
	for _, test := range tests {
		trie := Constructor()
		sut := &trie
		for _, FN := range test {
			FN(t, sut)
		}
	}
}
