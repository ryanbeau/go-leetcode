# Problem
[Leetcode - 124. Binary Tree Maximum Path Sum (Hard)](https://leetcode.com/problems/binary-tree-maximum-path-sum/)

## Takeaway
- Recursion will be the easier solution.
- We need the max value from the nodes following a path without passing through the same edge more than once.

## Solution
- Get the max value of the `Left` nodes, discard if less than zero.
- Get the max value of the `Right` nodes, discard if less than zero.
- If `Left` max + `Right` max + root node is higher than global `max`, set new max value.
- Return `root.Val` + either `Left` max or `Right` max depending on which is higher.
