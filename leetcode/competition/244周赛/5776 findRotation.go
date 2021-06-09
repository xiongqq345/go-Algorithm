package _44周赛

func findRotation(matrix [][]int, target [][]int) bool {
	n := len(matrix[0])
	var equal bool
	for i := 0; i < 4; i++ {
		equal = true
		for i := range matrix {
			for j := range matrix[0] {
				if matrix[i][j] != target[i][j] {
					equal = false
				}
			}
		}
		if equal {
			return true
		}

		for i := 0; i < n/2; i++ {
			matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
		}
		for i := 1; i < n; i++ {
			for j := 0; j < i; j++ {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
	return false
}
