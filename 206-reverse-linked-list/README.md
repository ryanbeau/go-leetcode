# Problem
[Leetcode - 206. Reverse Linked List (Easy)](https://leetcode.com/problems/reverse-linked-list/)

## Takeaway
- There are 2 different approaches.
  1. Iterative.
  2. Recursive.
- Iterative approach uses less resources than recursive and no stack is required.

## Solution
- Iterative approach reusing the nodes.
- Set current node's next to the current node and ensure the head's next is set to nil to prevent infinite loop.

Recursive solution:
```go
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	if head.Next != nil {
		current = reverseList(head.Next)
		head.Next.Next = head
	}
	head.Next = nil

	return current
}
```
