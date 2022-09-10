package solution

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//iterate to the nth node
	right := head
	for n > 0 {
		right = right.Next
		n--
	}
	if right == nil {
		return head.Next
	}

	//starting on nth node iterate until end
	left := head
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next
	return head
}
