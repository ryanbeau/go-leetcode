# Problem
[Leetcode - 100. Same Tree (Easy)](https://leetcode.com/problems/same-tree/)

## Takeaway
- Recursive will be very easy.

## Solution
- Use `isSameTree` to determine if both trees are not nil and then call `isSameTree` on their children.

The iterative approach is a bit more involved:
```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	nodes := []*TreeNode{p, q}
	for i := 0; i < len(nodes); i += 2 {
		if nodes[i].Val != nodes[i+1].Val {
			return false
		}
		if nodes[i].Left != nil && nodes[i+1].Left != nil {
			nodes = append(nodes, nodes[i].Left, nodes[i+1].Left)
		} else if nodes[i].Left != nil || nodes[i+1].Left != nil {
			return false
		}
		if nodes[i].Right != nil && nodes[i+1].Right != nil {
			nodes = append(nodes, nodes[i].Right, nodes[i+1].Right)
		} else if nodes[i].Right != nil || nodes[i+1].Right != nil {
			return false
		}
	}
	return true
}
```
