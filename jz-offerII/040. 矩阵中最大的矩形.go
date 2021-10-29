package jz_offerII

//给定一个由 0 和 1 组成的矩阵 matrix ，找出只包含 1 的最大矩形，并返回其面积。
//
//注意：此题 matrix 输入格式为一维 01 字符串数组。
func maximalRectangle(matrix []string) int {
	var ans int
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	for i := range matrix {
		for j, v := range matrix[i] {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}
	return ans
}
