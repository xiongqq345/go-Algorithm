package dfs_bfs

// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

func solve(grid [][]byte) {
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		dfs(grid, 0, i)
		dfs(grid, n-1, i)
	}
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 'A' {
				grid[i][j] = 'O'
			} else {
				grid[i][j] = 'X'
			}
		}
	}
}

func dfs(grid [][]byte, i, j int) {
	if inArea1(grid, i, j) && grid[i][j] == 'O' {
		grid[i][j] = 'A'
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
}

func inArea1(grid [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
