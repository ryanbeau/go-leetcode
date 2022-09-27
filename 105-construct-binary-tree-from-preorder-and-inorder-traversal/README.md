# Problem
[Leetcode - 105. Construct Binary Tree from Preorder and Inorder Traversal (Medium)](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

## Takeaway
- Recursive solution would be easier.
- First node in `preorder` is the root node.
- With the values in `inorder`, values left of `root`'s value are within `Left` node, and values right of `root`'s value are within the `Right` node.

## Solution
- Example `preorder: [3, 9, 20, 15, 7]` and `inorder: [9, 3, 15, 20, 7]`:
    ```
          3
         / \
        9   20
           /  \
          15   7
    ```

    Walkthrough of how to handle the recursion:
    ```
    preorder: [3, 9, 20, 15, 7]
    inorder:  [9, 3, 15, 20, 7]
    root val: 3
    Left:
        preorder: [9]
        inorder:  [9]
        root val: 9
    Right:
        preorder: [20, 15, 7]
        inorder:  [15, 20, 7]
        root val: 20
        Left:
            preorder: [15]
            inorder:  [15]
            root val: 15
        Right:
            preorder: [7]
            inorder:  [7]
            root val: 7
    ```
