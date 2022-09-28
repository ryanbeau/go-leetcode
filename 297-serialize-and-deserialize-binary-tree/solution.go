package solution

import (
	"strconv"
	"strings"
)

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	i := 0
	s := strings.Split(data, ",")

	var next func() *TreeNode
	next = func() *TreeNode {
		if s[i] != "" {
			v, _ := strconv.Atoi(s[i])
			i++
			return &TreeNode{
				Val:   v,
				Left:  next(),
				Right: next(),
			}
		}
		i++
		return nil
	}
	return next()
}
