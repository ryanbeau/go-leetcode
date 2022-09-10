# Problem
[Leetcode - 19. Remove Nth Node From End of List (Medium)](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## Takeaway
- We want to iterate through the nodes only once.
- To remove a node we need the left node.

## Solution
- Get the nth node from the left, and then iterate through with both head and nth node until we reach the end.
- A small shortcut is if the end node is nil then we know we need to remove the first node. (ie: array `[1,2,3]` with n `3` iterating to the nth node from the head would result in `2` -> `3` -> `nil` and `1` is the 3rd node from the end)
