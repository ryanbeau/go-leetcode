# Problem
[Leetcode - 121. Best Time to Buy and Sell Stock (Easy)](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

## Takeaway
- Sliding window solution.

## Solution
- Keep incrementing the right index while checking if `prices[l] < prices[r]`.
  - If true: Adjust max price if profit is higher.
  - If false: Adjust left index to equal right index because we found a new lower price.
