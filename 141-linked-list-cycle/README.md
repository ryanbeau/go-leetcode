# Problem
[Leetcode - 141. Linked List Cycle (Easy)](https://leetcode.com/problems/linked-list-cycle/)

## Takeaway
- We could use a map to store the ListNode as key.
- Or we could iterate through with two different speeds, if they land on each other then we have a cycle.

## Solution
- Going with the slow & fast iteration and comparison approach.
- Slow: increment by 1. Fast: increment by 2.
- If fast search finds nil then we don't have a cycle.
- If slow & fast are the same node we have a cycle.
