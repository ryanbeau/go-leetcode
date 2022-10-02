package solution

func maxProfit(prices []int) int {
	max, l := 0, 0
	for r := 1; r < len(prices); r++ {
		if prices[l] < prices[r] {
			if profit := prices[r] - prices[l]; profit > max {
				max = profit
			}
		} else {
			l = r
		}
	}
	return max
}
