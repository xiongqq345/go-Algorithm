package backtracking

// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var dfs func(vals []int, sum, index int)
	dfs = func(vals []int, sum, index int) {
		if sum > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int{}, vals...))
			return
		}

		for i := index; i < len(candidates); i++ {
			dfs(append(vals, candidates[i]), sum+candidates[i], i)
		}
	}
	dfs([]int{}, 0, 0)
	return ans
}
