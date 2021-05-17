package array

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	l, r, t, b := 0, n-1, 0, n-1
	num := 1
	for num <= n*n {
		for i := l; i <= r; i++ {
			matrix[t][i] = num
			num++
		}
		t++
		for i := t; i <= b; i++ {
			matrix[i][r] = num
			num++
		}
		r--
		for i := r; i >= l; i-- {
			matrix[b][i] = num
			num++
		}
		b--
		for i := b; i >= t; i-- {
			matrix[i][l] = num
			num++
		}
		l++
	}

	return matrix
}
