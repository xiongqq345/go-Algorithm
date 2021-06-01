package list_stack

// 给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
// 两个相邻元素间的距离为 1 。

var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

func updateMatrix(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	queue := make([][]int, 0)
	for i := range mat {
		for j := range mat[0] {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x, y := point[0]+dx[i], point[1]+dy[i]
			if x >= 0 && x < n && y >= 0 && y < m && mat[x][y] == -1 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}
	return mat
}
