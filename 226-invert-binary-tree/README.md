# Problem
[Leetcode - 226. Invert Binary Tree (Easy)](https://leetcode.com/problems/invert-binary-tree/)

## Takeaway
- Recursive will be very easy.

## Solution
- Use the `invertTree` function to swap both `Left` and `Right` if the parent is not `nil`.

The iterative solution is less simple:
```go
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Left != nil {
			nodes = append(nodes, nodes[i].Left)
		}
		if nodes[i].Right != nil {
			nodes = append(nodes, nodes[i].Right)
		}
		nodes[i].Left, nodes[i].Right = nodes[i].Right, nodes[i].Left
	}
	return root
}
```
