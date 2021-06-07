package greedy

// 有一个二维矩阵 A 其中每个元素的值为 0 或 1 。
//
//移动是指选择任一行或列，并转换该行或列中的每一个值：将所有 0 都更改为 1，将所有 1 都更改为 0。
//
//在做出任意次数的移动后，将该矩阵的每一行都按照二进制数来解释，矩阵的得分就是这些数字的总和。
//
//返回尽可能高的分数。
func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 1 << (n - 1) * m
	for j := 1; j < n; j++ {
		var sum int
		for _, row := range grid {
			if row[j] == row[0] {
				sum++
			}
		}
		if sum < m-sum {
			sum = m - sum
		}
		ans += 1 << (n - j - 1) * sum
	}
	return ans
}
