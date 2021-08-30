package backtracking

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	vis := make([]bool, n)
	var helper func([]int)
	helper = func(path []int) {
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			if !vis[i] {
				vis[i] = true
				helper(append(path, nums[i]))
				vis[i] = false
			}
		}
	}
	helper(nil)
	return ans
}
