package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type wordDictionaryFN func(t *testing.T, wd *WordDictionary)

func addWord(word string) wordDictionaryFN {
	return func(t *testing.T, wd *WordDictionary) {
		wd.AddWord(word)
	}
}

func assertSearch(word string, expected bool) wordDictionaryFN {
	return func(t *testing.T, wd *WordDictionary) {
		actual := wd.Search(word)
		assert.Equal(t, expected, actual, "Search(\"%s\")", word)
	}
}

func TestWordDictionary(t *testing.T) {
	tests := [][]wordDictionaryFN{
		{
			addWord("bad"),
			addWord("dad"),
			addWord("mad"),
			assertSearch("pad", false),
			assertSearch("bad", true),
			assertSearch(".ad", true),
			assertSearch("b..", true),
		},
	}
	for _, test := range tests {
		wd := Constructor()
		sut := &wd
		for _, FN := range test {
			FN(t, sut)
		}
	}
}
