package dfs_bfs

// 给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
//
//二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求：
//
//路径途经的所有单元格都的值都是 0 。
//路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
//畅通路径的长度 是该路径途经的单元格总数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shortest-path-in-binary-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func shortestPathBinaryMatrix(grid [][]int) int {
	rows := len(grid)
	if grid == nil || rows == 0 || grid[0][0] == 1 || grid[rows-1][rows-1] == 1 {
		return -1
	}
	if len(grid) == 1 && grid[0][0] == 0 {
		return 1
	}
	direction := []int{-1, 0, 1}
	grid[0][0] = 1
	//途经的每一个点都记录从起点到次的长度
	q := []int{0}
	//用que记录当前点的坐标，判断有没有下一个节点
	var x, y, nx, ny int
	for len(q) > 0 {
		x, y = q[0]/rows, q[0]%rows
		q = q[1:]
		for _, i := range direction {
			for _, j := range direction {
				if i == j && i == 0 {
					continue
				}
				nx, ny = x+i, y+j
				if nx < rows && ny < rows && nx >= 0 && ny >= 0 && grid[nx][ny] == 0 {
					q = append(q, nx*rows+ny)
					grid[nx][ny] = grid[x][y] + 1
					if nx == rows-1 && ny == rows-1 {
						return grid[nx][ny]
					}
				}
			}
		}
	}
	return -1
}
