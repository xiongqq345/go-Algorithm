package array

// 给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。
//todo
func countSmaller(nums []int) []int {
	ans := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		var cnt int
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[i] {
				cnt++
			}
		}
		ans[i] = cnt
	}
	return ans
}
