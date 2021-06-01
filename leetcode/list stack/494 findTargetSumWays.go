package list_stack

//给你一个整数数组 nums 和一个整数 target 。
//
//向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
//例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

func findTargetSumWays(nums []int, target int) int {
	var sum int
	var dfs func([]int, int)
	dfs = func(arr []int, cur int) {
		if len(arr) == 1 {
			if cur+arr[0] == target {
				sum++
			}
			if cur-arr[0] == target {
				sum++
			}
			return
		}
		dfs(arr[1:], cur+arr[0])
		dfs(arr[1:], cur-arr[0])
	}
	dfs(nums, 0)
	return sum
}
