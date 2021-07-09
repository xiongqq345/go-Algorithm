package dp

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, len(triangle[n-1]))
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	ans := math.MaxInt32
	for _, v := range dp {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func minimumTotal2(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < i+1; j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}
