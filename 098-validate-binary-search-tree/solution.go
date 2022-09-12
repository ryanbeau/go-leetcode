package solution

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValid(n *TreeNode, l int, r int) bool {
	return n == nil || n.Val < r && n.Val > l && isValid(n.Left, l, n.Val) && isValid(n.Right, n.Val, r)
}

func isValidBST(root *TreeNode) bool {
	return root == nil || isValid(root, math.MinInt, math.MaxInt)
}
