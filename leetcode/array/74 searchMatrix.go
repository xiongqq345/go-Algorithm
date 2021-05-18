package array

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	var row int
	for ; row < m-1; row++ {
		if target < matrix[row+1][0] {
			break
		}
	}

	for col := 0; col < n; col++ {
		if matrix[row][col] == target {
			return true
		}
	}
	return false
}
