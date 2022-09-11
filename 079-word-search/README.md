# Problem
[Leetcode - 79. Word Search (Medium)](https://leetcode.com/problems/word-search/)

## Takeaway
- Bruteforce method will be best where we look at each square on the board.

## Solution
- Start at each square of the `board` and determine if the `word` starts with that letter.
- If yes, then using recursion look at the north, south, east, west squares for the next letter and repeat.
  - Put down a temp placeholder on the square to ensure it doesn't repeat a previous letter.
