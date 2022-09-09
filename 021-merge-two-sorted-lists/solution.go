package solution

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		next := mergeTwoLists(list1.Next, list2)
		list1.Next = next
		return list1
	} else {
		next := mergeTwoLists(list1, list2.Next)
		list2.Next = next
		return list2
	}
}
