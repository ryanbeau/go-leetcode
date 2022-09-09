package solution

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	next := head.Next
	for next != nil {
		current, next, next.Next = next, next.Next, current
	}
	head.Next = nil

	return current
}
