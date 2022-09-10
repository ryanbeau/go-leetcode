# Problem
[Leetcode - 21. Merge Two Sorted Lists (Easy)](https://leetcode.com/problems/merge-two-sorted-lists/)

## Takeaway
- There are 2 different approaches.
  1. Iterative.
  2. Recursive.
- Iterative approach uses less resources than recursive but requires more logic.

## Solution
- Choosing the recursive approach reusing the nodes to reduce comprehension complexity.
- If the nodes are equal in value `list1` node should win.

The iterative approach as seen from [23. Merge k Sorted Lists - Solution](023-merge-k-sorted-lists).
```go
func mergeLists(list1 *ListNode, list2 *ListNode) *ListNode {
	temp := &ListNode{}
	last := temp

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			last.Next = list1
			list1 = list1.Next
		} else {
			last.Next = list2
			list2 = list2.Next
		}
		last = last.Next
	}
	if list1 != nil {
		last.Next = list1
	} else {
		last.Next = list2
	}
	return temp.Next
}
```
