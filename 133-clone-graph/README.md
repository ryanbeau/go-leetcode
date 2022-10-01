# Problem
[Leetcode - 133. Clone Graph (Medium)](https://leetcode.com/problems/clone-graph/)

## Takeaway
- A recursion solution will be the easiest.
- Create a new `Node` with the old value, and repeat for its neighbors.
- Neighbors will contain previously cloned nodes, so we need to ensure we're not cloning the same nodes ad infinitum.

## Solution
- Use a hash map for the old `Node` pointer with the value of the new `Node`.

1. If the old `Node` exists within the hash map, return the previously copied `Node`, otherwise create a new `Node` copy and add it to the hash map.
2. Iterate through the neighbors recursively while repeating from step 1.
