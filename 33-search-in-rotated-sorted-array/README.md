# Problem
[Leetcode - 33. Search in Rotated Sorted Array (Medium)](https://leetcode.com/problems/search-in-rotated-sorted-array/)

## Takeaway
- Binary search.
- Iterative approach instead of recursive.

## Solution
- If left half is low to high.
  - If `target` is within left half then progress search on left half.
  - Else progress search on right half.
- Else if left half is high to low.
  - If `target` is within right half then progress search on right half.
  - Else progress search on left half.
