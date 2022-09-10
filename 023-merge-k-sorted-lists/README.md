# Problem
[Leetcode - 23. Merge k Sorted Lists (Hard)](https://leetcode.com/problems/merge-k-sorted-lists/)

## Takeaway
- Using the [21. Merge Two Sorted Lists - Solution](021-merge-two-sorted-lists) to merge two ListNodes together.
- Keep iterating through the array until all are merged.

## Solution
- Rewrite the solution for merging two ListNodes as iterative instead of recursive.
- If the array has zero elements return `nil`.
- If the array has one element return the 1st.
- If the array has two or more elements, we can start with the 1st element as head ListNode and then merge it with each consecutive ListNode.
