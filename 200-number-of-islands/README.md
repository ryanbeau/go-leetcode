# Problem
[Leetcode - 200. Number of Islands (Medium)](https://leetcode.com/problems/number-of-islands/)

## Takeaway
- Recursion will be easiest.
- Track which cells in the grid have been visited before.
- Iterate through each cell in the grid looking for `'1'` and increment if it hasn't been visited before.

## Solution
- To keep track of the visited cells, we can use a single array with size of `totalRows*totalCols` and the cell index is: `row*totalCols + col`. The alternative is to change the cell's value in the `grid` array.

1. Iterate through each cell in the grid until we find `'1'` which indicates land.
    1. If it hasn't been visited before then increment the island count.
    2. Recursively visit each neighboring cell.
    3. Repeat.
2. Continue.
