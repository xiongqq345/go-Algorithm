package array

import "sort"

// 元素的 频数 是该元素在一个数组中出现的次数。
//
//给你一个整数数组 nums 和一个整数 k 。在一步操作中，你可以选择 nums 的一个下标，并将该下标对应元素的值增加 1 。
//
//执行最多 k 次操作后，返回数组中最高频元素的 最大可能频数 。
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	i, j, total := 0, 1, 0
	ans := 1
	for ; i < j && j < len(nums); j++ {
		total += (nums[j] - nums[j-1]) * (j - i)
		for total > k {
			total -= nums[j] - nums[i]
			i++
		}
		ans = max(ans, j-i+1)
	}
	return ans
}
