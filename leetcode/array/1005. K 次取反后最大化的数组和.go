package array

import "sort"

//给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
//
//选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
//重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
//
//以这种方式修改数组后，返回数组 可能的最大和 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	var sum int
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			k--
			nums[i] *= -1
		}
		sum += nums[i]
	}
	sort.Ints(nums)
	if k%2 == 0 {
		return sum
	}
	return sum - 2*nums[0]
}
