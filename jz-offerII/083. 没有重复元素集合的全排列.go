package jz_offerII

//给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	n := len(nums)
	vis := make([]bool, n)
	var ans [][]int
	var helper func(path []int)
	helper = func(path []int) {
		if len(path) == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for i := range nums {
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
