package 力扣杯竞赛

//小扣在秋日市集购买了一个古董键盘。由于古董键盘年久失修，键盘上只有 26 个字母 a~z 可以按下，且每个字母最多仅能被按 k 次。
//
//小扣随机按了 n 次按键，请返回小扣总共有可能按出多少种内容。由于数字较大，最终答案需要对 1000000007 (1e9 + 7) 取模。
//
func keyboard(k int, n int) int {
	comb := make([][]int, n+1)
	for i := 0; i < len(comb); i++ {
		comb[i] = make([]int, n+1)
	}
	comb[0][0] = 1
	for i := 1; i <= n; i++ {
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < 26; i++ {
		for j := n; j >= 1; j-- {
			for l := 1; l <= k; l++ {
				if l > j {
					break
				}
				dp[j] += dp[j-l] * comb[j][l]
				dp[j] %= 1e9 + 7
			}
		}
	}
	return dp[n]
}
