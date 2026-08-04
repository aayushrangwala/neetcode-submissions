func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	small := 0
	large := 1
	profit := 0

	for small < len(prices) && large < len(prices) {
		if prices[large] > prices[small] {
			profit = max(profit, prices[large] - prices[small])
			large++
			continue
		}

		small = large
		large++
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
