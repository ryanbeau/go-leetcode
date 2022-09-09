# Problem
[LeetCode - 271. Encode and Decode Strings (Medium)](https://leetcode.com/problems/encode-and-decode-strings/) 🔒 Locked

[LintCode - 659. Encode and Decode Strings (Medium)](https://www.lintcode.com/problem/659/) _available for free_

## Takeaway
- Using delimited strings would require escaping.
- Using a preamble to determine string length would be easiest.

## Solution
- Add a string length preamble where `["Hello","world"]` would encode as `"5)Hello5)world"`.
