package array

import "sort"

// 给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
func triangleNumber(nums []int) int {
	var cnt int
	sort.Ints(nums)
	for i := range nums {
		for j := i + 1; j < len(nums)-1; j++ {
			k := j + 1
			for ; k < len(nums) && nums[i]+nums[j] > nums[k]; k++ {
			}
			cnt += k - j - 1
		}
	}
	return cnt
}
