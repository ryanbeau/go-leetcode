# Problem
[Leetcode - 3. Longest Substring Without Repeating Characters (Easy)](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## Takeaway
- Sliding window solution.

## Solution
- Increment the right side while keeping track of the index position of each character.
  - If the current character exists within the window move the left of the window to the character's position plus 1.
  - Update the max substring size if the window length is greater.
