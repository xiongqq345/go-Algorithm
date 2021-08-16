package dp

// 给定一个01矩阵 M，找到矩阵中最长的连续1线段。这条线段可以是水平的、垂直的、对角线的或者反对角线的。
func longestLine(mat [][]int) int {
	var ans int
	n := len(mat[0])
	dp1 := make([]int, n+1)
	dp2 := make([]int, n)
	dp3 := make([]int, n+1)
	dp4 := make([]int, n+1)
	for i := range mat {
		var dp3ul int
		for j := range mat[0] {
			dp3tmp := dp3[j]
			if mat[i][j] == 1 {
				dp1[j+1] = dp1[j] + 1
				dp2[j] += 1
				dp3[j+1] = dp3ul + 1
				dp4[j] = dp4[j+1] + 1
				ans = max(ans, max(max(dp1[j+1], dp2[j]), max(dp3[j+1], dp4[j])))
			} else {
				dp1[j+1] = 0
				dp2[j] = 0
				dp3[j+1] = 0
				dp4[j] = 0
			}
			dp3ul = dp3tmp
		}
	}
	return ans
}
