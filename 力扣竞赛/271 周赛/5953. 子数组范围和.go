package _71_周赛

// 给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。
//
// 返回 nums 中 所有 子数组范围的 和 。
//
// 子数组是数组中一个连续 非空 的元素序列。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sum-of-subarray-ranges
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func subArrayRanges(nums []int) (ans int64) {
	for i, num := range nums {
		min, max := num, num
		for _, v := range nums[i+1:] {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
			ans += int64(max - min)
		}
	}
	return
}
