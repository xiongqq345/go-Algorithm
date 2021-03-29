package _21

import "math"

func maxProfit(prices []int) int {
	var (
		low       = math.MaxInt64
		maxProfit int
	)
	for _, price := range prices {
		low = min(low, price)
		maxProfit = max(maxProfit, price-low)
	}
	return maxProfit
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
