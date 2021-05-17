package array

//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
func rotate(matrix [][]int) {
	n := len(matrix[0])
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
