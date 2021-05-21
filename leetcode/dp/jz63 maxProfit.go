package dp

import "math"

//假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
func maxProfit(prices []int) int {
	buy, profit := math.MaxInt32, 0
	for _, price := range prices {
		buy = min(buy, price)
		profit = max(profit, price-buy)
	}
	return profit
}
