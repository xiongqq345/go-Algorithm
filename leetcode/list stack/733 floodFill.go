package list_stack

//有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。
//
//给你一个坐标(sr, sc)表示图像渲染开始的像素值（行 ，列）和一个新的颜色值newColor，让你重新上色这幅图像。
//
//为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。
//
//最后返回经过上色渲染后的图像。

var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs(image [][]int, row int, col int, originColor, newColor int) {
	if image[row][col] != originColor {
		return
	}
	image[row][col] = newColor
	for i := 0; i < 4; i++ {
		r, c := row+dx[i], col+dy[i]
		if r >= 0 && r < len(image) && c >= 0 && c < len(image[0]) {
			dfs(image, r, c, originColor, newColor)
		}
	}
}
