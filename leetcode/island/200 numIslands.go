package island

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。

func numIslands(grid [][]byte) int {
	var ans int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, i, j int) {
	if inArea1(grid, i, j) && grid[i][j] == '1' {
		grid[i][j] = '0'
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
}

func inArea1(grid [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
