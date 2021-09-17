package jz_offerII

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	var ans [][]int
	var helper func([]int, int)
	helper = func(path []int, start int) {
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for i := start; i <= n; i++ {
			helper(append(path, i), i+1)
		}
	}
	helper(nil, 1)
	return ans
}
