# Problem
[Leetcode - 230. Kth Smallest Element in a BST (Medium)](https://leetcode.com/problems/kth-smallest-element-in-a-bst/)

## Takeaway
- Iterative approach might be easiest.
- We need to navigate through the left-most nodes first.

## Solution
- Iterate through left most nodes first while adding them to a stack.
- When the current node is `nil` take from the top of the stack.

### Example: `root` (below) with `k` of `3`
```
         5
        / \
       3   6
      / \
     1   4
```

1. The stack will start with `[5]` then `[5,3]` then `[5,3,1]` and now there's no more `Left` nodes to add.
2. It then pops `[1]` off the stack, and increments the count to `1`.
3. There's no `Right` nodes with `[1]` so it continues by popping `[3]` off the stack and incrementing count to `2`.
4. It then adds `[4]` to the stack, and there are no `Left` nodes, so it increments count to `3`.
5. The count is now equal to `k` and it returns the value `4`.
