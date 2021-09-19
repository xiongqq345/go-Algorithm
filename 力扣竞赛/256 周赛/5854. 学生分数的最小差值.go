package _56_周赛

import "sort"

// 给你一个 下标从 0 开始 的整数数组 nums ，其中 nums[i] 表示第 i 名学生的分数。另给你一个整数 k 。
//
//从数组中选出任意 k 名学生的分数，使这 k 个分数间 最高分 和 最低分 的 差值 达到 最小化 。
//
//返回可能的 最小差值 。
//
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	var ans int = 1e6
	for i := k - 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-k+1])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
