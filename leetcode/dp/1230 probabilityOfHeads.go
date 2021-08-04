package dp

//有一些不规则的硬币。在这些硬币中，prob[i] 表示第 i 枚硬币正面朝上的概率。
//
//请对每一枚硬币抛掷 一次，然后返回正面朝上的硬币数等于 target 的概率。
//
func probabilityOfHeads(prob []float64, target int) float64 {
	n := len(prob)
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, target+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] * (1 - prob[i-1])
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j-1]*prob[i-1] + dp[i-1][j]*(1-prob[i-1])
		}
	}
	return dp[n][target]
}
