package solution

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max := root.Val
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lmax, rmax := 0, 0
		if l := dfs(node.Left); l > 0 {
			lmax = l
		}
		if r := dfs(node.Right); r > 0 {
			rmax = r
		}
		if m := node.Val + lmax + rmax; max < m {
			max = m
		}
		if lmax > rmax {
			return node.Val + lmax
		}
		return node.Val + rmax
	}
	dfs(root)
	return max
}
