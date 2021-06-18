package backtracking

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	var ans [][]int
	var helper func([]int, int)
	helper = func(vals []int, start int) {
		if len(vals) == k {
			ans = append(ans, append([]int{}, vals...))
			return
		}
		for i := start; i <= n; i++ {
			helper(append(vals, i), i+1)
		}
	}
	helper([]int{}, 1)
	return ans
}
