package array

// 给定一个非负整数数组，你最初位于数组的第一个位置。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
//假设你总是可以到达数组的最后一个位置。

// 从后向前贪心
func jump1(nums []int) int {
	pos := len(nums) - 1
	var step int
	for pos > 0 {
		for i := 0; i < pos; i++ {
			if nums[i]+i >= pos {
				pos = i
				step++
				break
			}
		}
	}
	return step
}

// 从前向后贪心
func jump(nums []int) int {
	var maxPos, end, step int
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == end {
			end = maxPos
			step++
		}
	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
