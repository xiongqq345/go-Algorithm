package coding

// 编写函数，实现许多图片编辑软件都支持的「颜色填充」功能。
//
//待填充的图像用二维数组 image 表示，元素为初始颜色值。初始坐标点的行坐标为 sr 列坐标为 sc。需要填充的新颜色为 newColor 。
//
//「周围区域」是指颜色相同且在上、下、左、右四个方向上存在相连情况的若干元素。
//
//请用新颜色填充初始坐标点的周围区域，并返回填充后的图像。
//
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	ori := image[sr][sc]
	if image[sr][sc] == newColor {
		return image
	}
	var helper func(int, int)
	helper = func(i, j int) {
		if inArea(image, i, j) && image[i][j] == ori {
			image[i][j] = newColor
			helper(i+1, j)
			helper(i-1, j)
			helper(i, j+1)
			helper(i, j-1)
		}
	}
	helper(sr, sc)
	return image
}

func inArea(grid [][]int, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
