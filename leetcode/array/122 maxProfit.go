package array

//给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
func maxProfit2(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
