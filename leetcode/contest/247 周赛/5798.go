package _47_周赛

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	var tmp int
	for layer := 0; layer < min(m, n)/2; layer++ {
		kk := (m + n - 2 - layer*4) * 2
		for ; kk > 0; kk-- {
			l, r, t, b := layer, n-1-layer, layer, m-1-layer
			for i := t; i <= b; i++ {
				tmp, grid[i][l] = grid[i][l], tmp
			}
			l++
			for i := l; i <= r; i++ {
				tmp, grid[b][i] = grid[b][i], tmp
			}
			b--
			for i := b; i >= t; i-- {
				tmp, grid[i][r] = grid[i][r], tmp
			}
			r--
			for i := r; i >= l; i-- {
				tmp, grid[t][i] = grid[t][i], tmp
			}
			t++
			grid[layer][layer] = tmp
		}
	}

	return grid
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
