package dfs_bfs

//给你一个整数数组 nums 和一个整数 target 。
//
//向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
//例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

func findTargetSumWays(nums []int, target int) int {
	var cnt int
	var helper func([]int, int)
	helper = func(arr []int, cur int) {
		if len(arr) == 1 {
			if cur+arr[0] == target {
				cnt++
			}
			if cur-arr[0] == target {
				cnt++
			}
			return
		}
		helper(arr[1:], cur+arr[0])
		helper(arr[1:], cur-arr[0])
	}
	helper(nums, 0)
	return cnt
}
