# Problem
[Leetcode - 104. Maximum Depth of Binary Tree (Easy)](https://leetcode.com/problems/maximum-depth-of-binary-tree/)

## Takeaway
- Recursive will be very easy.

## Solution
- Use the `maxDepth` function to get the max value of `Left` and `Right` plus 1 if the parent is not `nil`, otherwise return 0.

The iterative approach is a bit more involved:
```go
func maxDepth(root *TreeNode) (max int) {
	if root == nil {
		return 0
	}
	nodes, depth := []*TreeNode{root}, []int{1}
	for i := 0; i < len(nodes); i++ {
		if max < depth[i] {
			max = depth[i]
		}
		if nodes[i].Left != nil {
			nodes = append(nodes, nodes[i].Left)
			depth = append(depth, depth[i]+1)
		}
		if nodes[i].Right != nil {
			nodes = append(nodes, nodes[i].Right)
			depth = append(depth, depth[i]+1)
		}
	}
	return max
}
```
