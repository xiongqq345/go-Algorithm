package bfs

// 在给定的网格中，每个单元格可以有以下三个值之一：
//
//值 0 代表空单元格；
//值 1 代表新鲜橘子；
//值 2 代表腐烂的橘子。
//每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
//
//返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。
//
func orangesRotting(grid [][]int) int {
	res, orange := 0, 0 // 时间,好橘子个数
	var queue []int     // 怀橘子队列
	n := len(grid[0])
	// 遍历二维数组,将烂橘子的坐标加入队列
	for i := range grid {
		for j := range grid[0] {
			switch grid[i][j] {
			case 1:
				orange++
			case 2:
				queue = append(queue, i*n+j)
			}
		}
	}
	// 上下左右四个方向
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	// bfs
	for orange > 0 && len(queue) != 0 {
		size := len(queue) // 提前保存队列长度
		for i := 0; i < size; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			c, r := node%n, node/n
			// 怀橘子的上下左右遍历
			for j := 0; j < 4; j++ {
				nr, nc := r+dx[j], c+dy[j]
				if inArea(grid, nr, nc) && grid[nr][nc] == 1 {
					orange--         // 好橘子减一
					grid[nr][nc] = 2 // 好橘子变成坏橘子,然后入队
					queue = append(queue, nr*n+nc)
				}
			}
		}
		// 时间加一
		res++
	}
	if orange != 0 { // 好橘子有剩余
		return -1
	}
	return res
}

func inArea(grid [][]int, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
