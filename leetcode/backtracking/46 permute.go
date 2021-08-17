package backtracking

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	vis := make([]bool, n)
	var dfs func([]int)
	dfs = func(vals []int) {
		if len(vals) == n {
			ans = append(ans, append([]int{}, vals...))
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			vis[i] = true
			dfs(append(vals, nums[i]))
			vis[i] = false
		}
	}
	dfs([]int{})
	return ans
}
