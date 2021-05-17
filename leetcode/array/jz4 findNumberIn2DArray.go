package array

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	row, col := 0, cols-1
	for row < rows && col >= 0 {
		switch {
		case target > matrix[row][col]:
			row++
		case target < matrix[row][col]:
			col--
		default:
			return true
		}
	}
	return false
}
