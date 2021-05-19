package array

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
func minPathSum(grid [][]int) int {
	n := len(grid[0])
	dp := make([]int, n)
	for i := range grid {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}
