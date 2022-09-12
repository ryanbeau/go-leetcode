package solution

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		if l, r := maxDepth(root.Left), maxDepth(root.Right); l > r {
			return l + 1
		} else {
			return r + 1
		}
	}
	return 0
}
