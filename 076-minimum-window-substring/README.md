# Problem
[Leetcode - 76. Minimum Window Substring (Hard)](https://leetcode.com/problems/minimum-window-substring/)

## Takeaway
- A sliding window approach would be best for an O(N) complexity.

## Solution
- Build a hashmap containing the count of each character of `t` and the count contained within the window.
- Increment the right index when all characters are not contained in the window.
- Increment the left index when all characters are contained in the window.
