package _8_双周赛

import "sort"

// 你正在设计一个动态数组。给你一个下标从 0 开始的整数数组 nums ，其中 nums[i] 是 i 时刻数组中的元素数目。除此以外，你还有一个整数 k ，表示你可以 调整 数组大小的 最多 次数（每次都可以调整成 任意 大小）。
//
//t 时刻数组的大小 sizet 必须大于等于 nums[t] ，因为数组需要有足够的空间容纳所有元素。t 时刻 浪费的空间 为 sizet - nums[t] ，总 浪费空间为满足 0 <= t < nums.length 的每一个时刻 t 浪费的空间 之和 。
//
//在调整数组大小不超过 k 次的前提下，请你返回 最小总浪费空间 。
//
//注意：数组最开始时可以为 任意大小 ，且 不计入 调整大小的操作次数。

func minSpaceWastedKResizing(nums []int, k int) int {
	var tmp []int
	for i := 0; i < len(nums)-1; i++ {
		tmp = append(tmp, abs(nums[i+1]-nums[i]))
	}
	if len(tmp) == 0 {
		return 0
	}
	sort.Ints(tmp)
	var ans int
	for i := range tmp[:len(tmp)-k] {
		ans += tmp[i]
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
