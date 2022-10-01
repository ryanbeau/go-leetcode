# Problem
[Leetcode - 417. Pacific Atlantic Water Flow (Medium)](https://leetcode.com/problems/pacific-atlantic-water-flow/)

## Takeaway
- A recursion solution will be easiest.
- We want to avoid the brute force approach of iterating through every cell in the height grid.
- Start on the edges and with depth first search check if the cell adjacent to it is higher.
- Keep record of which cells have been visited from each edge, Pacific or Atlantic.

## Solution
- Have a hash map that contains the visited cells.
    - If the map's value is 1 it got here from Pacific.
    - If the map's value is 2 it got here from Atlantic.
    - If the ocean's value differs from the current, add the cell to the results.
        - Set the value as negative to prevent duplicate results.
- For Pacific dfs start from the cells on the North and West edges.
- For Atlantic dfs start from the cells on the South and East edges.

1. Validate if the cell is higher or equal in height to the previous cell.
2. Ensure the cell hasn't been visited before from the same ocean.
3. Finally we recursively check the North, East, South, & West cells starting back at #1.
