package solution

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeLists(list1 *ListNode, list2 *ListNode) *ListNode {
	temp := &ListNode{}
	end := temp

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			end.Next = list1
			list1 = list1.Next
		} else {
			end.Next = list2
			list2 = list2.Next
		}
		end = end.Next
	}
	if list1 != nil {
		end.Next = list1
	} else {
		end.Next = list2
	}
	return temp.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	head := lists[0]
	for i := 1; i < len(lists); i++ {
		if lists[i] != nil {
			head = mergeLists(head, lists[i])
		}
	}

	return head
}
