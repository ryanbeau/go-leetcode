package solution

//Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	copied := make(map[*Node]*Node, 4)

	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if copy, has := copied[node]; has {
			return copy
		}
		copy := &Node{Val: node.Val}
		copied[node] = copy
		for _, neighbor := range node.Neighbors {
			copy.Neighbors = append(copy.Neighbors, deepCopy(neighbor))
		}
		return copy
	}
	return deepCopy(node)
}
