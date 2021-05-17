package array

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	l, r, t, b := 0, n-1, 0, m-1
	ans := make([]int, m*n)
	num := 0
	for num < m*n {
		for i := l; i <= r; i++ {
			ans[num] = matrix[t][i]
			num++
		}
		t++
		for i := t; i <= b; i++ {
			ans[num] = matrix[i][r]
			num++
		}
		r--
		if num == m*n {
			break
		}
		for i := r; i >= l; i-- {
			ans[num] = matrix[b][i]
			num++
		}
		b--
		for i := b; i >= t; i-- {
			ans[num] = matrix[i][l]
			num++
		}
		l++
	}
	return ans
}
