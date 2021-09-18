package jz_offerII

//给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
//
//candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
//
//对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
//
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var helper func([]int, int, int)
	helper = func(path []int, pos, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for i := pos; i < len(candidates); i++ {
			helper(append(path, candidates[i]), i, sum+candidates[i])
		}
	}
	helper(nil, 0, 0)
	return ans
}
