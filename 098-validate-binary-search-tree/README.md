# Problem
[Leetcode - 98. Validate Binary Search Tree (Medium)](https://leetcode.com/problems/validate-binary-search-tree/)

## Takeaway
- Recursion will be the easiest approach.
- With the constraints of `Val` we can take some shortcuts on comparing outermost `Left` and `Right` nodes, `Val` will never reach the min/max range of `int`.
  >- `-2³¹` <= `Node.val` <= `2³¹ - 1`

## Solution
- When checking any descendant nodes within a `Left` node we should ensure it's less than the previous `Val` and not greater than the ancestor `Val`.
- When checking any descendant nodes within a `Right` node we should ensure it's greater than the previous `Val` and not less than the ancestor `Val`.
