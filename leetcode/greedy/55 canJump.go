package greedy

// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个下标。
func canJump(nums []int) bool {
	n := len(nums)
	var maxPos int
	for i, v := range nums {
		if i <= maxPos {
			maxPos = max(maxPos, i+v)
			if maxPos >= n-1 {
				return true
			}
		} else {
			break
		}
	}
	return false
}
