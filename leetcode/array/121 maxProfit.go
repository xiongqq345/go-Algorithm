package array

func maxProfit(prices []int) int {
	buy, profit := prices[0], 0
	for _, price := range prices {
		buy = min(buy, price)
		profit = max(profit, price-buy)
	}
	return profit
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
