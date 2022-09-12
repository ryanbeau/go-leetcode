# Problem
[Leetcode - 102. Binary Tree Level Order Traversal (Medium)](https://leetcode.com/problems/binary-tree-level-order-traversal/)

## Takeaway
- Iterative approach might be easiest.
- Append value level by level.

## Solution
- Start on root level, iterating level by level adding the `Val`, then add `Left` or `Right` if not `nil` for the next level of nodes.
