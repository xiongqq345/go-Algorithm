package dfs_bfs

// 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
//
//对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func longestIncreasingPath(matrix [][]int) int {
	ds := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	var helper func(int, int) int
	helper = func(i, j int) int {
		if dp[i][j] != 0 {
			return dp[i][j]
		}

		dp[i][j]++
		for _, d := range ds {
			r, c := i+d[0], j+d[1]
			if inArea(matrix, r, c) && matrix[r][c] > matrix[i][j] {
				dp[i][j] = max(dp[i][j], helper(r, c)+1)
			}
		}
		return dp[i][j]
	}

	var ans int
	for i := range matrix {
		for j := range matrix[0] {
			ans = max(ans, helper(i, j))
		}
	}
	return ans
}
