package jz_offer

// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
func movingCount(m int, n int, k int) int {
	var count int
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var helper func(int, int)
	helper = func(i int, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}

		if numSum(i)+numSum(j) > k {
			return
		}
		vis[i][j] = true
		count++
		helper(i+1, j)
		helper(i-1, j)
		helper(i, j+1)
		helper(i, j-1)
		vis[i][j] = false
	}
	helper(0, 0)
	return count
}

func numSum(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
