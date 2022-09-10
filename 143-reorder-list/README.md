# Problem
[Leetcode - 143. Reorder List (Medium)](https://leetcode.com/problems/reorder-list/)

## Takeaway
- Iterative approach.
- Interweave the first half with the last half.
- With the example input of `[1,2,3,4,5,6]`
  - The first half is found at the head node `[1,2,3]`.
  - The last half will need to be retrieved and reversed `[6,5,4]`.
  - Interweaving the two halves to create `[1,6,2,5,3,4]`.

## Solution
- Find the middle node. This can be found by iterating through the ListNode with a single and double lookahead and middle is found when the double lookahead reaches the end.
- Reverse the last half of the nodes.
- Merge the left and right halves of the ListNodes.
