package solution

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	nodes, results := []*TreeNode{root}, [][]int{}
	for i, d := 0, 1; i < d; d = len(nodes) {
		var order []int
		for ; i < d; i++ {
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
			order = append(order, nodes[i].Val)
		}
		results = append(results, order)
	}
	return results
}
