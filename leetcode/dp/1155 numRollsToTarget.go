package dp

//这里有 d 个一样的骰子，每个骰子上都有 f 个面，分别标号为 1, 2, ..., f。
//
//我们约定：掷骰子的得到总点数为各骰子面朝上的数字的总和。
//
//如果需要掷出的总点数为 target，请你计算出有多少种不同的组合情况（所有的组合情况总共有 f^d 种），模 10^9 + 7 后返回。
func numRollsToTarget(d int, f int, target int) int {
	dp := make([][]int, d+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	for i := 1; i <= d; i++ {
		for j := 1; j <= target; j++ {
			for k := 1; k <= f && j >= k; k++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % (1e9 + 7)
			}
		}
	}
	return dp[d][target]
}
