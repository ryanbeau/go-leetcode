package solution

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// get middle node
	middle, temp := head, head.Next
	for temp != nil && temp.Next != nil {
		middle = middle.Next
		temp = temp.Next.Next
	}

	// reverse second half
	var last *ListNode
	temp, middle.Next = middle.Next, nil
	for temp != nil {
		last, temp, temp.Next = temp, temp.Next, last
	}

	// merge first & last halves
	first := head
	for last != nil {
		first, last, first.Next, last.Next = first.Next, last.Next, last, first.Next
	}
}
