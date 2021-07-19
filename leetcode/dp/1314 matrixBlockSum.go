package dp

// 给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素 mat[r][c] 的和：
//
//i - k <= r <= i + k,
//j - k <= c <= j + k 且
//(r, c) 在矩阵内。
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	P := make([][]int, m+1)
	for i := range P {
		P[i] = make([]int, n+1)
	}
	// P是二维前缀和数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			P[i][j] = P[i-1][j] + P[i][j-1] - P[i-1][j-1] + mat[i-1][j-1]
		}
	}

	for i := range ans {
		for j := range ans[0] {
			ans[i][j] = get(P, m, n, i+k+1, j+k+1) + get(P, m, n, i-k, j-k) - get(P, m, n, i+k+1, j-k) - get(P, m, n, i-k, j+k+1)
		}
	}
	return ans
}

func get(pre [][]int, m, n, x, y int) int {
	x = max(min(x, m), 0)
	y = max(min(y, n), 0)
	return pre[x][y]
}
