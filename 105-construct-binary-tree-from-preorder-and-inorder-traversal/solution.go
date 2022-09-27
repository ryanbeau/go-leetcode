package solution

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			if i > 0 {
				root.Left = buildTree(preorder[1:i+1], inorder[:i])
			}
			if i < len(inorder)-1 {
				root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			}
			break
		}
	}
	return root
}
