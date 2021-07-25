package coding

// 硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
func waysToChange(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{25, 10, 5, 1}
	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[n] % (1e9 + 7)
}
