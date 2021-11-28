package _69_周赛

import "sort"

//给你一个下标从 0 开始的整数数组 nums 以及一个目标元素 target 。
//
//目标下标 是一个满足 nums[i] == target 的下标 i 。
//
//将 nums 按 非递减 顺序排序后，返回由 nums 中目标下标组成的列表。如果不存在目标下标，返回一个 空 列表。返回的列表必须按 递增 顺序排列。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-target-indices-after-sorting-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	var ans []int
	for i := sort.SearchInts(nums, target); i < len(nums); i++ {
		if nums[i] != target {
			break
		}
		ans = append(ans, i)
	}
	return ans
}
