# Problem
[Leetcode - 235. Lowest Common Ancestor of a Binary Search Tree (Medium)](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)

## Takeaway
- Recursive approach will be the easiest.
- We can take some shortcuts due to:
  >- All `Node.val` are unique.
  >- `p` and `q` will exist in the BST.

## Solution
- Shortcuts due to node values being unique and `p` and `q` must exist.
  - If `p.Val` and `q.val` are less than `root.Val`, continue recursion to `root.Left`.
  - If `p.Val` and `q.val` are greater than `root.Val`, continue recursion to `root.Right`.
  - Otherwise return `root`.
