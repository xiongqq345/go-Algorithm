package island

// 给定一个包含了一些 0 和 1 的非空二维数组 grid 。
//
//一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
//
//找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
func maxAreaOfIsland(grid [][]int) int {
	var maxArea int
	for i := range grid {
		for j := range grid[0] {
			var area int
			dfsMaxArea(grid, i, j, &area)
			if maxArea < area {
				maxArea = area
			}
		}
	}
	return maxArea
}

func dfsMaxArea(grid [][]int, i, j int, area *int) {
	if inArea(grid, i, j) && grid[i][j] == 1 {
		grid[i][j] = 2
		*area++
		dfsMaxArea(grid, i+1, j, area)
		dfsMaxArea(grid, i-1, j, area)
		dfsMaxArea(grid, i, j+1, area)
		dfsMaxArea(grid, i, j-1, area)
	}
}

func inArea(grid [][]int, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
