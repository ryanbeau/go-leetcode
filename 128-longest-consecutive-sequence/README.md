# Problem
[LeetCode - 128. Longest Consecutive Sequence (Hard)](https://leetcode.com/problems/longest-consecutive-sequence/)

## Takeaway
- Sorting the array would be the simple approach but it would not run in `O(n)` time.
- Hashmap would be the optimal approach.

## Solution
- Iterate through the items and increment with the count from the leftmost and rightmost sequence for the hashmap value.
