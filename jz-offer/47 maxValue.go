package jz_offer

// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
func maxValue(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	for _, row := range grid {
		for j, cur := range row {
			if j == 0 {
				dp[j] += cur
			} else {
				dp[j] = max(dp[j-1], dp[j]) + cur
			}
		}
	}
	return dp[len(dp)-1]
}
