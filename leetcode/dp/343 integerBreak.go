package dp

import "math"

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		var imax int
		for j := 1; j < i; j++ {
			imax = max(imax, max(j*dp[i-j], j*(i-j)))
		}
		dp[i] = imax
	}
	return dp[n]
}

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	quo, rem := n/3, n%3
	if rem == 0 {
		return int(math.Pow(3, float64(quo)))
	} else if rem == 1 {
		return int(math.Pow(3, float64(quo-1))) * 4
	}
	return int(math.Pow(3, float64(quo))) * 2
}
