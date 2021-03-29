package _74

func searchMatrix(matrix [][]int, target int) bool {
	var col int
	for i := 0; i < len(matrix); i++ {
		if target < matrix[i][0] {
			col = i - 1
			break
		}
		if i == len(matrix)-1 {
			col = i
		}
	}
	if col < 0 {
		return false
	}

	for i := 0; i < len(matrix[col]); i++ {
		if matrix[col][i] == target {
			return true
		}
	}
	return false
}
