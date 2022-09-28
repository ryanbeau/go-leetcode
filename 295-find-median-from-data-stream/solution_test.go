package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type finderFN func(t *testing.T, finder *MedianFinder)

func add(num int) finderFN {
	return func(t *testing.T, finder *MedianFinder) {
		finder.AddNum(num)
	}
}

func assertFindMedian(expected float64) finderFN {
	return func(t *testing.T, finder *MedianFinder) {
		actual := finder.FindMedian()
		assert.Equal(t, expected, actual, "FindMedian()")
	}
}

func TestMedianFinder(t *testing.T) {
	tests := [][]finderFN{
		{
			add(1),
			add(5),
			add(2),
			add(4),
			add(3),
			assertFindMedian(3.0),
		},
		{
			add(2),
			add(3),
			add(4),
			assertFindMedian(3.0),
		},
		{
			add(4),
			add(3),
			add(2),
			assertFindMedian(3.0),
		},
		{
			add(2),
			add(3),
			assertFindMedian(2.5),
		},
		{
			add(3),
			add(2),
			assertFindMedian(2.5),
		},
		{
			add(1),
			add(2),
			assertFindMedian(1.5),
			add(3),
			assertFindMedian(2.0),
		},
		{
			add(2),
			assertFindMedian(2.0),
		},
	}
	for i, test := range tests {
		sut := Constructor()
		t.Run(fmt.Sprintf("Test#%d", i+1), func(t *testing.T) {
			for _, FN := range test {
				FN(t, &sut)
			}
		})
	}
}
