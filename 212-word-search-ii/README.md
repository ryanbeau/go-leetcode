# Problem
[Leetcode - 212. Word Search II (Hard)](https://leetcode.com/problems/word-search-ii/)

## Takeaway
- This is similar to [79. Word Search - Solution](079-word-search) but with multiple words.
- Bruteforce method will be best where we look at each square on the board.
- With the constraints of lowercase english letters we can safely assume only `'a'` to `'z'`.
  > - `board[i][j]` is a lowercase English letter.
  > - `words[i]` consists of lowercase English letters.

## Solution
- Use a Trie data structure for each of the words.
- Start at each square of the `board` and determine if any of the `words` starts with that letter (if the Trie has children of that character).
- If yes, then using recursion look at the north, south, east, west squares for the next letter and repeat.
  - Put down a temp placeholder on the square to ensure it doesn't repeat a previous letter.
- Prevent the found `words` within the Trie from being used more than once.
