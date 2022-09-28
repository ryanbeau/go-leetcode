# Problem
[Leetcode - 297. Serialize and Deserialize Binary Tree (Hard)](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)

## Takeaway
- The problem asks for the output of serialize to differ from how Leetcode serializes a binary tree.
- Recursion is easiest for this problem.
- Comma delimited is fine for serialization and `nil` nodes can be empty values.

## Solution
- Serialize the data such as that the root then leftmost nodes followed by rightmost nodes are output in order. Example:
  ```
        3
       / \
      9   20
         /  \
        15   7
  ```
  Should serialize to: `"3,9,,,20,15,,,7,,"`
- Deserialize should recursively get the current node followed by `Left` nodes then `Right` nodes in order.
