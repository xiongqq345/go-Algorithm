package _4__双周赛

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 1
	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			maxLen := min(m-i, n-j)
			for k := 2; k <= maxLen; k++ {
				flag := true
				sum := sumSubRowArr(grid[i][j : k+j])
				for r := i + 1; r < i+k; r++ {
					if sumSubRowArr(grid[r][j:k+j]) != sum {
						flag = false
						break
					}
				}
				if !flag {
					break
				}
				for c := j; c < j+k; c++ {
					if sumSubColArr(grid, c, i, i+k-1) != sum {
						break
					}
				}
				if !flag {
					continue
				}

				var sum1 int
				for r, c := i, j; r < i+k; {
					sum1 += grid[r][c]
					r++
					c++
				}
				if sum1 != sum {
					flag = false
				}

				var sum2 int
				for r, c := i, j; r < i+k; {
					sum2 += grid[i+k-r][j+k-c]
					r++
					c++
				}
				if sum2 != sum {
					flag = false
				}

				if flag {
					ans = max(ans, k)
				}
			}
		}
	}
	return ans
}

func sumSubRowArr(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func sumSubColArr(arr [][]int, col, start, end int) int {
	var sum int
	for i := start; i <= end; i++ {
		sum += arr[i][col]
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
