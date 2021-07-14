package dp

// 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
//
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// s1: 手上持有股票的最大收益
	// s2: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// s3: 手上不持有股票，并且不在冷冻期中的累计最大收益
	var s1, s2, s3 int
	s1 = -prices[0]
	for i := 1; i < n; i++ {
		tmp1, tmp2, tmp3 := s1, s2, s3
		s1 = max(tmp1, tmp3-prices[i])
		s2 = tmp1 + prices[i]
		s3 = max(tmp2, tmp3)
	}
	return max(s2, s3)
}
