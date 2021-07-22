package dp

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n)
	var maxSide int
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[j], maxSide = 1, 1
		}
	}
	var upLeft int
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp := dp[j]
			if j == 0 && matrix[i][j] == '1' {
				dp[j] = 1
			} else {
				if matrix[i][j] == '1' {
					dp[j] = min(dp[j], min(dp[j-1], upLeft)) + 1
				} else {
					dp[j] = 0
				}
			}
			upLeft = tmp
			maxSide = max(maxSide, dp[j])
		}
	}
	return maxSide * maxSide
}
