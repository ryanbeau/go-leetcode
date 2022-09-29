# Problem
[Leetcode - 39. Combination Sum (Medium)](https://leetcode.com/problems/combination-sum/)

## Takeaway
- Recursion will be the easier solution.
- We will want to iterate through each element in the `candidates` array even multiple times.
    - For instance with `candidates` of `[1,2,3]` and `target` of `3` we want these results:
        - `[1,1,1]`
        - `[1,1,2]`
        - `[3]`

## Solution
- Start on the first indices of `candidates` (`i=0`)
    1. If the cumulative value is less than `target` then recursively branch to adding itself again.
    2. Then recursively branch to adding the next `candidates` (`i+1`).
    3. If the cumulative value is equal to `target` then add the vals to `results`.
    4. Repeat.
