package coding

func rotate(matrix [][]int) {
	m := len(matrix)
	for i := 0; i < m/2; i++ {
		matrix[i], matrix[m-i-1] = matrix[m-i-1], matrix[i]
	}
	for i := range matrix {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
