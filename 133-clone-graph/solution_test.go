package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	var n1, n2, n3, n4 *Node = &Node{Val: 1}, &Node{Val: 2}, &Node{Val: 3}, &Node{Val: 4}
	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}
	/* Makes the following Node:
	 *    1 -- 2
	 *    |    |
	 *    4 -- 3
	 */

	for _, input := range []*Node{n1, n2, n3, n4} {
		checked := make(map[*Node]struct{})

		//a custom assertion method since assert.NotSame does not iterate through fields
		var assertEqual func(expected *Node, actual *Node)
		assertEqual = func(expected *Node, actual *Node) {
			if _, has := checked[actual]; !has {
				assert.NotSame(t, expected, actual)
				assert.NotSame(t, expected.Neighbors, actual.Neighbors)

				assert.EqualValues(t, expected.Val, actual.Val)
				assert.Len(t, actual.Neighbors, len(expected.Neighbors))

				checked[actual] = struct{}{}
				for i := 0; i < len(actual.Neighbors); i++ {
					assertEqual(expected.Neighbors[i], actual.Neighbors[i])
				}
			}
		}

		actual := cloneGraph(input)
		assertEqual(input, actual)
	}
}
