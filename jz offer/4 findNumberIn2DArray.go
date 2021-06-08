package jz_offer

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	r, c := m-1, 0
	for r >= 0 && c < n {
		switch {
		case matrix[r][c] > target:
			r--
		case matrix[r][c] < target:
			c++
		default:
			return true
		}
	}
	return false
}
