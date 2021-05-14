package array

//找出数组中重复的数字。
//
//
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = struct{}{}
	}
	return 0
}

func findRepeatNumber2(nums []int) int {
	set := make([]bool, len(nums))
	for _, v := range nums {
		if set[v] {
			return v
		}
		set[v] = true
	}
	return 0
}
