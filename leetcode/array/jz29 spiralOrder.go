package array

// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	l, t, r, b := 0, 0, n-1, m-1
	ans, num := make([]int, m*n), 0
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
